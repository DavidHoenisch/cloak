package execs

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func getHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}

	return home
}

func Test_splitCommandOnSpace(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []string
		wantErr bool
	}{
		{
			name: "Get parent command with no args",
			args: args{
				cmd: "ls",
			},
			want:    "ls",
			want1:   nil,
			wantErr: false,
		},
		{
			name: "Get parent command and args",
			args: args{
				cmd: "ls -la",
			},
			want:    "ls",
			want1:   []string{"-la"},
			wantErr: false,
		},
		{
			name: "attempt command split with '' command provided",
			args: args{
				cmd: "",
			},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := splitCommandOnSpace(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitCommandOnSpace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("splitCommandOnSpace() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("splitCommandOnSpace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRunner_ExecCommandInNewProcess(t *testing.T) {
	type args struct {
		c     string
		group string
	}
	tests := []struct {
		name    string
		r       *Runner
		args    args
		shell   string
		wantErr bool
		envPath string
	}{
		{
			name: "Run ls command in subprocess",
			args: args{
				c:     "ls -la",
				group: "",
			},
			wantErr: false,
			envPath: "./test_env.json",
			shell:   "/bin/bash",
		},
		{
			name: "Run command with expected error",
			args: args{
				c:     "ls -la",
				group: "",
			},
			wantErr: false,
			envPath: "./test_env.json",
			shell:   "/bin/fish",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{}
			if err := r.ExecCommandInNewProcess(tt.args.c, tt.args.group, tt.envPath, tt.shell); (err != nil) != tt.wantErr {
				t.Errorf("Runner.ExecCommandInNewProcess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunner_getGroupEnvVars(t *testing.T) {
	type args struct {
		group   string
		envPath string
	}
	tests := []struct {
		name string
		r    *Runner
		args args
		want []string
	}{
		{
			name: "ensure vars formatted as key=val",
			args: args{
				group:   "ExampleGroup",
				envPath: fmt.Sprintf("%s/.cloak/env.json", getHome()),
			},
			want: []string{
				"ANTHROPIC_API_KEY=some-random-string",
				"OPENAI_API_KEY=some-random-string",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{}
			if got := r.getGroupEnvVars(tt.args.group, tt.args.envPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Runner.getGroupEnvVars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunner_parseCommandString(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		r    *Runner
		args args
		want *CommandParts
	}{
		{
			name: "test get nil *CommandParts struct",
			args: args{
				command: "",
			},
			want: nil,
		},
		{
			name: "test get nil *CommandParts struct",
			args: args{
				command: "ls -la",
			},
			want: &CommandParts{
				Command: "ls",
				Args:    []string{"-la"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{}
			if got := r.parseCommandString(tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Runner.parseCommandString() = %v, want %v", got, tt.want)
			}
		})
	}
}
