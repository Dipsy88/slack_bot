package plugins

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"slack-bot/extra"
	"strings"
	"time"

	"github.com/go-chat-bot/bot"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func cloneAdd(command *bot.Cmd) (msg string, err error) {
	//input := command.Args

	text := command.RawArgs
	input := strings.Split(text, " ")

	if len(input) < 2 {
		fmt.Println("Passed arguments", len(input))
		msg = "Pass the git repo and path to clone to"
	} else {
		var config extra.Configuration
		if len(input) == 3 {
			config = extra.GetConfig(filepath.FromSlash(input[2]))
		} else {
			config = extra.GetConfig("")
		}

		path := filepath.FromSlash(input[1])
		makeDirIfRequired(path)

		repo, err := git.PlainClone(path, false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: config.GitUser,
				Password: config.GitPass,
			},
			URL:      input[0],
			Progress: os.Stdout,
		})
		checkIfError(err)

		w, err := repo.Worktree()
		checkIfError(err)

		filename := filepath.Join(path, "ExampleFile")
		err = ioutil.WriteFile(filename, []byte("hello world {}"), 0644)
		checkIfError(err)

		// Add the file to the staging area
		_, err = w.Add("ExampleFile")
		checkIfError(err)

		// Verify the current status of the worktree using the method Status
		status, err := w.Status()
		checkIfError(err)

		fmt.Println(status)

		// Check commit
		commit, err := w.Commit("ExampleFile", &git.CommitOptions{
			Author: &object.Signature{
				Name:  config.GitUser,
				Email: config.GitEmail,
				When:  time.Now(),
			},
		})
		checkIfError(err)

		obj, err := repo.CommitObject(commit)
		checkIfError(err)
		fmt.Println(obj)

		err = repo.Push(&git.PushOptions{
			Auth: &http.BasicAuth{
				Username: config.GitUser,
				Password: config.GitPass,
			},
		})
		checkIfError(err)

		msg = input[0]
	}
	return msg, nil
}

func init() {
	bot.RegisterCommand(
		"clone", "git clone to a destination",
		"path",
		cloneAdd)
}

func makeDirIfRequired(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		fmt.Println("Creating directory", path)
		err := os.MkdirAll(path, 0755)
		checkIfError(err)
	}
}

func checkIfError(e error) {
	if e != nil {
		panic(e)
	}
}
