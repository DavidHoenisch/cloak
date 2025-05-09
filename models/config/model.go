package config

type Config struct {
	EnvPath   string `json:"env_path,omitempty"`
	ShellPath string `json:"shell_path,omitempty"`
}
