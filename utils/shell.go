package utils

import (
	"errors"
	"os"
)

func GetUserShell() (string, error) {
	shell := os.Getenv("SHELL")

	if shell == "" {
		return shell, errors.New("shell not set in env")
	}

	return shell, nil
}
