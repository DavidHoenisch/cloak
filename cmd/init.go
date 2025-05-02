/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"cloak/internal/config"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new cloak config",
	Run: func(cmd *cobra.Command, args []string) {
		force := cmd.Flags().Changed("force")

		err := config.GenerateConfigFile(force, "")
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "overwrite existing config")
}
