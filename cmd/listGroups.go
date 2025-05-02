/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listGroupsCmd represents the listGroups command
var listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "list all the environment groups defined in your cloak config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listGroups called")
	},
}

func init() {
	rootCmd.AddCommand(listGroupsCmd)

	// listGroupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
