package main

import (
	"os"
	_ "slack-bot/plugins"

	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/chucknorris"
)

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
