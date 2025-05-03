/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate the cloak config file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validate called")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
