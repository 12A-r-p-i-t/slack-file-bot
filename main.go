package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5405759576979-5418528455713-HlcHe9EWM0CMxW3sSGP0F6vK")
	os.Setenv("CHANNEL_ID", "C05BXNBHALT")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{}

	for i := 0; i < len(fileArr); i++ {

	}
}
