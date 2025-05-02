/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "a collection of commands for working with the cloak config command",
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(initCmd)
	configCmd.AddCommand(validateCmd)
	configCmd.AddCommand(listGroupsCmd)
}
