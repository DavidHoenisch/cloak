/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"cloak/internal/execs"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// cmdCmd represents the cmd command
var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "run a command with cloak environmental vars",
	Long: `Run a command with cloaked environmental vars.

For best results, quote out the entire commend string`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the default value for the env flag after Settings is initialized
		if !cmd.Flags().Changed("env") {
			env = Settings.DefaultEnvPath
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		group := cmd.Flag("group").Value.String()
		command := cmd.Flag("command").Value.String()

		fmt.Println(env)

		r := execs.Runner{}

		err := r.ExecCommandInNewProcess(command, group, env)

		if err != nil {
			log.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(cmdCmd)

	cmdCmd.Flags().StringP("command", "c", "", "command to run")
	cmdCmd.Flags().StringP("group", "g", "", "group environment to inject")
	cmdCmd.Flags().StringVarP(&env, "env", "e", "", "path to env file. Leave blank for default")

	// mark command and group as required args
	cmdCmd.MarkFlagRequired("command")
	cmdCmd.MarkFlagRequired("group")

}
