package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

// slack token(This value setted at build as option -ldflag)
var token string

func main() {
	// Init slack client
	api := slack.New(token)

	var message string

	// Parse commandline args
	argMessage := flag.String("message", "", "Text to post to Slack")
	channel := flag.String("channel", "random", "Text of the channel to be posted. Default value is random")
	username := flag.String("username", "Notifier", "Text of UserName who post notification")
	flag.Parse()

	// Check message value
	if *argMessage == "" {
		// If message text from arg is null, read from stdin
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		message = stdin.Text()
		if message == "" {
			fmt.Println("Empty message text. Please try again.")
			return
		}
	} else {
		message = *argMessage
	}

	// Posting message
	_, _, err := api.PostMessage("#"+*channel, message, slack.PostMessageParameters{Username: *username})
	if err != nil {
		fmt.Println("Faild to post message")
		fmt.Println(err)
	}
	return
}
