/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version string

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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
