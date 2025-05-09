// package defines default configs for cloak

package settings

import (
	"fmt"
	"log/slog"
	"os"
)

type Settings struct {
	// defaults to $HOME/.cloak/env.json
	DefaultEnvPath       string
	DefaultEnvParentPath string
	// defaults to $HOME/.config/cloak/conf.json
	DefaultConfigPath       string
	DefaultConfigParentPath string

	// Shell related commands
	SystemShell string
}

// folder and file names
const (
	XdgFolder       string = "cloak"
	EnvHiddenFolder string = ".cloak"
	ConfigFileName  string = "conf.json"
	envFileName     string = "env.json"
)

// env var names
const (
	CLOAK_ENV_PATH    string = "CLOAK_ENV_PATH"
	CLOAK_CONFIG_PATH string = "CLOAK_CONFIG_PATH"
)

var (
	// full path to the default config
	fullConfigPath string = fmt.Sprintf("%s/%s/%s", getUserConfigDirectory(),
		XdgFolder,
		ConfigFileName)

	// path the parent folder
	fullConfigFolderPath string = fmt.Sprintf("%s/%s", getUserConfigDirectory(),
		XdgFolder)

	// path the users directory
	fullHomePath string = getUserHomeDirectory()

	// Path to the env directory
	fullEnvPath string = fmt.Sprintf("%s/%s/%s", getUserHomeDirectory(),
		EnvHiddenFolder,
		envFileName)

	// path to the env folder
	fullEnvFolderPath string = fmt.Sprintf("%s/%s", getUserHomeDirectory(),
		EnvHiddenFolder)
)

// Gets the XDG home values from the environmental vars
func getUserConfigDirectory() string {
	configHome, err := os.UserConfigDir()
	if err != nil {
		slog.Error("could not find user home dir")
	}

	return configHome
}

// gets the users home director
func getUserHomeDirectory() string {
	userHome, err := os.UserHomeDir()
	if err != nil {
		slog.Error("could not find user home dir")
	}

	return userHome
}

//TODO: create the ability to configure settings with env vars
// func getSettingsFromEnv() []string {
//
//
//
// 	return nil
// }

func New() *Settings {
	return &Settings{
		DefaultEnvPath:          fullEnvPath,
		DefaultEnvParentPath:    fullEnvFolderPath,
		DefaultConfigPath:       fullConfigPath,
		DefaultConfigParentPath: fullConfigFolderPath,
	}

}
