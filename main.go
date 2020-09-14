package main

import (
	"fmt"
	"os"
	_ "slack-bot/plugins"

	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/chucknorris"
)

func main() {
	fmt.Println("Slack token is", os.Getenv("SLACK_TOKEN"))
}
