package config

import (
	"cloak/models/config"
	"encoding/json"
	"log"
	"os"
)

func readOutFileAsBytes() []byte {
	d, err := os.ReadFile(fullConfigPath)
	if err != nil {
		log.Println(err)
	}

	return d
}

// parses config file, if errors are encountered during unmarshalling
// and error is returned
func ParseInConfigFile(configPath string) (*config.Config, error) {

	var conf config.Config
	return &conf, json.Unmarshal(readOutFileAsBytes(), &conf)

}
