package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
)

// slack token(This value setted at build as option -ldflag)
var token string

func main() {
	// Init slack client
	api := slack.New(token)

	// Parse commandline args
	message := flag.String("message", "", "Text to post to Slack")
	channel := flag.String("channel", "random", "Text of the channel to be posted. Default value is random")
	flag.Parse()

	// Posting message
	_, _, err := api.PostMessage("#"+*channel, *message, slack.PostMessageParameters{Username: "Notifier"})
	if err != nil {
		fmt.Println("Faild to post message")
		fmt.Println(err)
	}
	return
}
