/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"github.com/DavidHoenisch/cloak/utils"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listGroupsCmd represents the listGroups command
var listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "list all the environment groups defined in your cloak config",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the default value for the env flag after Settings is initialized
		if !cmd.Flags().Changed("path") {
			env = Settings.DefaultEnvPath
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		configPath := cmd.Flag("path").Value.String()

		groups, err := utils.GetVarGroups(configPath)
		if err != nil {
			log.Println(err)
		}

		for _, v := range groups {
			fmt.Println(v)
		}

	},
}

func init() {
	listGroupsCmd.Flags().StringVarP(&env, "path", "p", "", "custom path to file")
}
