package extra

import (
	"fmt"
	"log"

	"github.com/tkanos/gonfig"
)

var (
	filePath string = "/Users/Dipesh/Documents/pass/configuration.json"
)

// Configuration stores the config
type Configuration struct {
	GitUser  string
	GitEmail string
	GitPass  string
}

// GetConfig exports the configuration
func GetConfig(path string) Configuration {
	if path == "" {
		path = filePath
	}
	fmt.Println(path)
	configuration := Configuration{}
	err := gonfig.GetConf(filePath, &configuration)
	if err != nil {
		log.Println(err)
	}
	return configuration
}
