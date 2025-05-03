package config

import (
	"cloak/models/env"
	"encoding/json"
	"log"
	"os"
)

func readOutFileAsBytes(path string) []byte {
	d, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	return d
}

// parses config file, if errors are encountered during unmarshalling
// and error is returned
func ParseInConfigFile(configPath string) (*env.Env, error) {

	var conf env.Env
	return &conf, json.Unmarshal(readOutFileAsBytes(configPath), &conf)

}
