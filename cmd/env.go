/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"github.com/DavidHoenisch/cloak/internal/config"
	"github.com/DavidHoenisch/cloak/models/types"
	"log"

	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "create a new env file ($HOME/.cloak/env.json)",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the default value for the env flag after Settings is initialized
		if !cmd.Flags().Changed("path") {
			env = Settings.DefaultEnvPath
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		force := cmd.Flags().Changed("force")
		customPath := cmd.Flag("path").Value.String()

		err := config.GenerateFile(force, customPath, types.Env)
		if err != nil {
			log.Println(err)
		}

	},
}

func init() {

	envCmd.Flags().BoolP("force", "f", false, "overwrite existing config")
	envCmd.Flags().StringVarP(&env, "path", "p", "", "custom path to file")
}
