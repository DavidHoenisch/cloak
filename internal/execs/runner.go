package execs

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"cloak/internal/config"
)

type CommandParts struct {
	Command string
	Args    []string
}

type RunCommandWithEnvs interface {
	parseCommandString(command string) (string, []string)
	getGroupEnvVars(group string) []string
	ExecCommandInNewProcess(group string) error
}

type Runner struct{}

// parseCommandString() will take in the while command
// and will do a "best effort" attempt at splitting the command
// up into the app being called and its respective arguments
func (r *Runner) parseCommandString(command string) *CommandParts {

	c, args, err := splitCommandOnSpace(command)
	if err != nil {
		return nil
	}

	return &CommandParts{
		Command: c,
		Args:    args,
	}
}

func (r *Runner) getGroupEnvVars(group, envPath string) []string {

	var vars []string

	grps, err := config.ParseInConfigFile(envPath)
	if err != nil {
		return nil
	}

	groups := grps.Groups

	for _, v := range groups {
		if v.Name == group {
			for _, vv := range v.Vars {
				readyVar := fmt.Sprintf("%s=%s", vv.Key, vv.Value)
				vars = append(vars, readyVar)

			}
		}
	}

	return vars
}

func (r *Runner) ExecCommandInNewProcess(c, group, envPath string) error {
	ctx := context.TODO()

	command := r.parseCommandString(c)
	cmd := exec.CommandContext(ctx, command.Command, command.Args...)

	for _, v := range r.getGroupEnvVars(group, envPath) {
		cmd.Env = append(cmd.Environ(), v)
	}

	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(out))

	return nil
}

// this function handles taking in an entire command sequence
// and splitting it up into the parent command and its arguments
// allowing it to be passed into the exec.Cmd method
func splitCommandOnSpace(cmd string) (string, []string, error) {
	var command string
	var args []string

	pCommand := strings.Split(cmd, " ")
	if len(pCommand) >= 1 {
		command = pCommand[0]

		// create the slice of arguments
		for i, v := range pCommand {
			if i == 0 {
				continue
			} else {
				args = append(args, v)
			}
		}

	} else {
		command = cmd
	}

	return command, args, nil
}
