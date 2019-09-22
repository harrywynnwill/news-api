package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var settings map[string]interface{}

func GetSettings() map[string]interface{} {
	configPath := "./config.json"
	file, e := ioutil.ReadFile(configPath)
	if e != nil {
		log.Fatalf("Error loading config.json: %v\n", e)
		os.Exit(1)
	}

	if e := json.Unmarshal(file, &settings); e != nil {
		log.Fatalf("Error reading config.json: %v\n", e)
		os.Exit(1)
	}

	return settings
}
