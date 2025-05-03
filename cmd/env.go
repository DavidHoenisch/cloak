/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cloak/internal/config"
	"cloak/models/types"
	"log"

	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the default value for the env flag after Settings is initialized
		if !cmd.Flags().Changed("path") {
			env = Settings.DefaultEnvPath
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		force := cmd.Flags().Changed("force")
		customPath := cmd.Flag("path").Value.String()

		err := config.GenerateEnvFile(force, customPath, types.Config)
		if err != nil {
			log.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(envCmd)

	envCmd.Flags().BoolP("force", "f", false, "overwrite existing config")
	envCmd.Flags().StringP("path", "p", "", "custom path to file")
}
