/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
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
	Short: "create a new cloak config file ($HOME/.config/cloak/conf.json)",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the default value for the env flag after Settings is initialized
		if !cmd.Flags().Changed("path") {
			env = Settings.DefaultConfigPath
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

	confCmd.Flags().BoolP("force", "f", false, "overwrite existing config")
	confCmd.Flags().StringP("path", "p", "", "custom path to file")
}
