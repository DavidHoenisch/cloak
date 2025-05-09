package utils

import (
	"reflect"
	"testing"
)

func TestGetVarGroups(t *testing.T) {
	type args struct {
		configPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test getting all groups out of config file",
			args: args{
				configPath: "./test_env.json",
			},
			want:    []string{"ExampleGroup"},
			wantErr: false,
		},
		{
			name: "test when bad config path is provided",
			args: args{
				configPath: "./nopath.json",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVarGroups(tt.args.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVarGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVarGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
