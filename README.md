# go-notifier
Post message to slack.  
This command is useful for notify of long task ended.  
Before build, you need Slack API Token.

## Build
Before build, you need add package `github.com/nlopes/slack`.

```
$ SLACK_TOKEN=xxxxxxxxxxxxxxxxxxx make
```

## Usage
```
$ notifier --message "text of posting message" --channel "text of channel name"
```

`--message` option is required. `--channel` option is optional.
