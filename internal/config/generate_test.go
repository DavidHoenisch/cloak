package config

import (
	"testing"

	"github.com/DavidHoenisch/cloak/models/types"
)

func Test_createDefaultDirectory(t *testing.T) {
	type args struct {
		t types.Ftype
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test create new when one already exists",
			args: args{
				t: types.Config,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createDefaultDirectory(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("createDefaultDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkConfigIfExists(t *testing.T) {
	type args struct {
		t types.Ftype
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkConfigIfExists(tt.args.t); got != tt.want {
				t.Errorf("checkConfigIfExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createFile(t *testing.T) {
	type args struct {
		t types.Ftype
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createFile(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("createFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateFile(t *testing.T) {
	type args struct {
		force       bool
		custom      string
		configOrEnv types.Ftype
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateFile(tt.args.force, tt.args.custom, tt.args.configOrEnv); (err != nil) != tt.wantErr {
				t.Errorf("GenerateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
