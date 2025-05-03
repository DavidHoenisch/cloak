/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"cloak/utils"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listGroupsCmd represents the listGroups command
var listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "list all the environment groups defined in your cloak config",
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
	rootCmd.AddCommand(listGroupsCmd)

	// listGroupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listGroupsCmd.Flags().StringP("path", "p", Settings.DefaultEnvPath, "custom path to file")
}
