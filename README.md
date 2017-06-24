# TONKO2🎉

Google Play Application first release 🎉 Checker in Go

Please use [tonkotsu](https://github.com/operando/tonkotsu) to check for updates.


## Run

```
go run tonko2.go config.go -c config.toml
```

## Config File


```toml
log = "debug"
sleeptime = 1
error_post = true

[slack_update_post]
text = "Update!!!!!"
username = "TONKOTSU bot"
icon_emoji = ":pig:"
channel = "#test"
link_names = true

[slack_error_post]
text = "Error!!!!!"
username = "bot"
icon_emoji = ":ghost:"
channel = "#test"
link_names = true

[slack_start_post]
text = "Running tonkotsu..."
username = "bot"
icon_emoji = ":ghost:"
channel = "#test"
link_names = true

[webhook]
url = "webhook_url" # your Incoming WebHooks URL for Slack

[android]
package = "com.mercariapp.mercari" # your Android application package name
```

## License

```
Copyright 2017 Shinobu Okano

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```