package config

import (
	"cloak/internal/settings"
	"cloak/models/config"
	"cloak/models/env"
	"cloak/models/types"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
)

const (
	XdgFolder       string = "cloak"
	EnvHiddenFolder string = ".cloak"
	ConfigFileName  string = "conf.json"
	envFileName     string = "env.json"
)

var Settings settings.Settings = *settings.New()

var (
	// example config to use creating the initial file
	exampleEnv env.Env = env.Env{
		Name: "EnvFileName",
		Groups: []env.Group{
			{
				Name: "ExampleGroup",
				Vars: []env.KeyValue{
					{
						Key:   "ANTHROPIC_API_KEY",
						Value: "some-random-string",
					},
					{
						Key:   "OPENAI_API_KEY",
						Value: "some-random-string",
					},
				},
			},
		},
	}

	exampleConfig config.Config = config.Config{
		EnvPath: "some path to somewhere",
	}
)

// Creates the cloak config file in the users XDG home
func createDefaultDirectory(t types.Ftype) error {

	switch t {
	case types.Config:
		return os.Mkdir(Settings.DefaultConfigParentPath, 0755)
	case types.Env:
		return os.Mkdir(Settings.DefaultEnvParentPath, 0755)
	default:
		return errors.New("error figuring out whether to create a config or env dir")
	}
}

// checkConfigIfExists() returns true if cannot find
// config file. Returns false if no error encountered finding
// config file
func checkConfigIfExists(t types.Ftype) bool {

	switch t {
	case types.Config:
		_, statErr := os.Stat(Settings.DefaultConfigPath)
		if os.IsNotExist(statErr) {
			return false
		} else {
			return true
		}
	case types.Env:
		_, statErr := os.Stat(Settings.DefaultEnvPath)
		if os.IsNotExist(statErr) {
			return false
		} else {
			return true
		}
	default:
		return true
	}
}

func createFile(t types.Ftype) error {

	switch t {
	case types.Config:
		var err error
		err = createDefaultDirectory(t)

		content, err := json.MarshalIndent(&exampleConfig, "", "	")

		err = os.WriteFile(Settings.DefaultConfigPath,
			content,
			0644)

		if err != nil {
			slog.Error("error encountered while creating file")
		}
		return err

	case types.Env:

		var err error
		err = createDefaultDirectory(t)

		content, err := json.MarshalIndent(&exampleEnv, "", "	")

		err = os.WriteFile(Settings.DefaultEnvPath,
			content,
			0644)

		if err != nil {
			slog.Error("error encountered while creating file")
		}
		return err

	}

	return nil
}

// Generates config file if not exists
//
//	force: allows overwriting of existing config
//	custom: custom path for the config file
func GenerateFile(force bool, custom string, configOrEnv types.Ftype) error {

	switch checkConfigIfExists(configOrEnv) {
	case false:
		return createFile(configOrEnv)

	case true:
		switch force {
		case true:
			return createFile(configOrEnv)
		case false:
			return errors.New("file already exist. Use --force if you intent to overwrite")
		default:
			return nil
		}
	default:
		return nil
	}
}
