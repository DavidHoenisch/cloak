package execs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/DavidHoenisch/cloak/internal/config"
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

// substituteEnvVars replaces $VAR and ${VAR} patterns in the command string
// with actual values from the provided environment variables
func (r *Runner) substituteEnvVars(command string, envVars []string) string {
	// Create a map of environment variables for quick lookup
	envMap := make(map[string]string)
	for _, env := range envVars {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	// Regular expression to match $VAR or ${VAR} patterns
	re := regexp.MustCompile(`\$\{([^}]+)\}|\$([A-Za-z_][A-Za-z0-9_]*)`)

	result := re.ReplaceAllStringFunc(command, func(match string) string {
		var varName string
		if strings.HasPrefix(match, "${") {
			// Handle ${VAR} format
			varName = match[2 : len(match)-1]
		} else {
			// Handle $VAR format
			varName = match[1:]
		}

		if value, exists := envMap[varName]; exists {
			return value
		}
		// If variable not found, return the original match
		return match
	})

	return result
}

func (r *Runner) ExecCommandInNewProcess(c, group, envPath, shell string) error {
	// NOTE: The context may need to be passed in from elsewhere
	ctx := context.TODO()

	// Get group environment variables
	groupEnvVars := r.getGroupEnvVars(group, envPath)

	// Substitute environment variables in the command string
	substitutedCommand := r.substituteEnvVars(c, groupEnvVars)

	// command := r.parseCommandString(c)
	// cmd := exec.CommandContext(ctx, command.Command, command.Args...)
	cmd := exec.CommandContext(ctx, shell, "-c", substitutedCommand)

	// Start with the current environment
	cmd.Env = cmd.Environ()

	// Append all group environment variables
	for _, v := range groupEnvVars {
		cmd.Env = append(cmd.Env, v)
	}

	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return err
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

	// check if command is blank
	if cmd == "" {
		return "", nil, errors.New("blank command provided")
	}

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
