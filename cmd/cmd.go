/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cloak/internal/execs"
	"log"

	"github.com/spf13/cobra"
)

// cmdCmd represents the cmd command
var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "run a command with cloak environmental vars",
	Long: `Run a command with cloaked environmental vars.

For best results, quote out the entire commend string`,
	Run: func(cmd *cobra.Command, args []string) {
		group := cmd.Flag("group").Value.String()
		command := cmd.Flag("command").Value.String()

		r := execs.Runner{}

		err := r.ExecCommandInNewProcess(command, group)

		if err != nil {
			log.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(cmdCmd)

	cmdCmd.Flags().StringP("command", "c", "", "command to run")
	cmdCmd.Flags().StringP("group", "g", "", "group environment to inject")

	// mark command and group as required args
	cmdCmd.MarkFlagRequired("command")
	cmdCmd.MarkFlagRequired("group")

}
