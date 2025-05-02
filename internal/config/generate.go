package config

import (
	"cloak/models/config"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

const (
	XdgFolder      string = "cloak"
	ConfigFileName string = "conf.json"
)

var (
	// full path to the default config
	fullConfigPath string = fmt.Sprintf("%s/%s/%s", getUserConfigDirectory(),
		XdgFolder,
		ConfigFileName)

	fullConfigFolderPath string = fmt.Sprintf("%s/%s", getUserConfigDirectory(),
		XdgFolder)
)

// Creates the cloak config file in the users XDG home
func createDefaultConfigDirectory() error {
	return os.Mkdir(fullConfigFolderPath, 0755)

}

// Gets the XDG home values from the environmental vars
func getUserConfigDirectory() string {
	configHome, err := os.UserConfigDir()
	if err != nil {
		slog.Error("could not find user home dir")
	}

	return configHome
}

// checkConfigIfExists() returns true if cannot find
// config file. Returns false if no error encountered finding
// config file
func checkConfigIfExists() bool {
	_, statErr := os.Stat(fullConfigPath)

	return errors.Is(statErr, os.ErrExist)
}

func createConfig() error {
	var err error
	err = createDefaultConfigDirectory()

	content, err := json.MarshalIndent(&config.Config{}, "", "	")

	err = os.WriteFile(fullConfigPath,
		content,
		0644)

	if err != nil {
		slog.Error("error encountered while creating config",
			err)
	}
	return err
}

// Generates config file if not exists
//
//	force: allows overwriting of existing config
//	custom: custom path for the config file
func GenerateConfigFile(force bool, custom string) error {

	switch checkConfigIfExists() {
	// this would mean errors encountered finding config
	case false:
		return createConfig()

	case true:
		switch force {
		case true:
			return createConfig()
		case false:
			return errors.New("config file already exist. Use --force if you intent to overwrite")
		default:
			return nil
		}
	default:
		return nil
	}
}
