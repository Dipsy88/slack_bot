package main

import (
	"fmt"
	"net/http"
	"slack-bot/extra"
	_ "slack-bot/plugins"
	"sync"

	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/chucknorris"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		http.HandleFunc("/", handler)
		http.ListenAndServe(":3000", nil)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var config extra.Configuration
		config = extra.GetConfig("")
		slack.Run(config.SlackToken)
		//slack.Run(os.Getenv("SLACK_TOKEN"))
	}()
	wg.Wait()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congratulations! Your Go application has been successfully deployed on Kubernetes.")
}
