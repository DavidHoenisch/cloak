// package defines default configs for cloak

package settings

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

func Test_getUserConfigDirectory(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test get config directory",
			want: fmt.Sprintf("%s/.config", getHome()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserConfigDirectory(); got != tt.want {
				t.Errorf("getUserConfigDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUserHomeDirectory(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test get home directory",
			want: getHome(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserHomeDirectory(); got != tt.want {
				t.Errorf("getUserHomeDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Settings
	}{
		{
			name: "test get settings",
			want: &Settings{
				DefaultEnvPath:          fmt.Sprintf("%s/.cloak/env.json", getHome()),
				DefaultEnvParentPath:    fmt.Sprintf("%s/.cloak", getHome()),
				DefaultConfigPath:       fmt.Sprintf("%s/.config/cloak/conf.json", getHome()),
				DefaultConfigParentPath: fmt.Sprintf("%s/.config/cloak", getHome()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
