/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"os"

	"github.com/DavidHoenisch/cloak/internal/settings"

	"github.com/spf13/cobra"
)

func getDefaultConfigs() settings.Settings {
	return *settings.New()
}

var Version string = "dev-build"

// Global settings for the app
var Settings settings.Settings

var env string
var shell string
var conf string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: Version,
	Use:     "cloak",
	Short:   "segment environmental vars into groups and only expose to apps explicity",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() { Settings = getDefaultConfigs() })
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
