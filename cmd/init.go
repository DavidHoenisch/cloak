/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new cloak config or env",
}

func init() {
	initCmd.AddCommand(envCmd)
	initCmd.AddCommand(confCmd)
}
