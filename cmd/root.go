/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"cloak/internal/settings"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string
var Settings settings.Settings

func getDefaultConfigs() settings.Settings {
	Sts := settings.New()
	return *Sts
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: version,
	Use:     "cloak",
	Short:   "segment environmental vars into groups and only expose to cli apps explicity",
	Long:    ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		Settings = getDefaultConfigs()
		fmt.Println("Initialized settings ", Settings)
	})
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
