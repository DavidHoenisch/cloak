package execs

import (
	"reflect"
	"testing"
)

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
		wantErr bool
	}{
		{
			name: "Run ls command in subprocess",
			args: args{
				c:     "ls -la",
				group: "",
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{}
			if err := r.ExecCommandInNewProcess(tt.args.c, tt.args.group); (err != nil) != tt.wantErr {
				t.Errorf("Runner.ExecCommandInNewProcess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
