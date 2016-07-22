
package main

import (
	"flag"
	"fmt"
	"time"
    "net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/operando/golack"
)

const (
	GOOGLE_PLAY = "https://play.google.com/store/apps/details?id="
)

func createGooglePlayURL(android Android) string {
	googlePlayURL := GOOGLE_PLAY + android.Package
	log.Debug(googlePlayURL)
	return googlePlayURL
}

func main() {
	var configPath = flag.String("c", "", "configuration file path")
	flag.Parse()

	var config Config
	_, err := LoadConfig(*configPath, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	setLogLevel(config.Log)
	sleep := time.Duration(config.SleepTime*60) * time.Second

	var googlePlayURL string
	uPayload := golack.Payload{
		config.SlackUpdatePost,
	}

	ePayload := golack.Payload{
		config.SlackErrorPost,
	}

	if config.SlackStartPost.Text != "" {
		sPayload := golack.Payload{
			config.SlackStartPost,
		}
		golack.Post(sPayload, config.Webhook)
	}

	if config.Android.Package == "" {
		log.Debug("Package is empty.")
	} else {
		googlePlayURL = createGooglePlayURL(config.Android)
		log.Info("Check Google Play URL : " + googlePlayURL)
	}

	log.Info("Slack Post Message : " + config.SlackUpdatePost.Text)
	log.Info("Slack Errro Message : " + config.SlackErrorPost.Text)

	for {
        resp, err := http.Get(googlePlayURL)
        log.Info(resp)
        if err != nil && config.ErrorPost {
            ePayload.Slack.Text = ePayload.Slack.Text + "\n" + err.Error()
            log.Error("Slack Errro Message : " + ePayload.Slack.Text)
            golack.Post(ePayload, config.Webhook)
        } else {
            if resp.StatusCode != 404 {
                golack.Post(uPayload, config.Webhook)
                log.Info("Release!!!!!!!!!!!")
                break
            } else {
                log.Info("No Release")
            }
        }
		time.Sleep(sleep)
	}

	log.Info("Release check end.")
}

func init() {
	log.SetLevel(log.InfoLevel)
}

func setLogLevel(lv string) {
	switch lv {
	case "debug", "d":
		log.SetLevel(log.DebugLevel)
	case "info", "i":
		log.SetLevel(log.InfoLevel)
	case "warn", "w":
		log.SetLevel(log.WarnLevel)
	case "error", "e":
		log.SetLevel(log.ErrorLevel)
	case "fatal", "f":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
