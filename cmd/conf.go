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

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	rootCmd.AddCommand(confCmd)

	confCmd.Flags().BoolP("force", "f", false, "overwrite existing config")
	confCmd.Flags().StringP("path", "p", Settings.DefaultConfigPath, "custom path to file")
}
