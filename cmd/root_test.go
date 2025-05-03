/*
Copyright © 2025 DavidHoenisch dh1689@pm.me
*/
package cmd

import (
	"cloak/internal/settings"
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

func Test_getDefaultConfigs(t *testing.T) {
	tests := []struct {
		name string
		want settings.Settings
	}{
		{
			name: "test get populated settings",
			want: settings.Settings{
				DefaultEnvPath:          fmt.Sprintf("%s/.cloak/env.json", getHome()),
				DefaultEnvParentPath:    fmt.Sprintf("%s/.cloak", getHome()),
				DefaultConfigPath:       fmt.Sprintf("%s/.config/cloak/conf.json", getHome()),
				DefaultConfigParentPath: fmt.Sprintf("%s/.config/cloak", getHome()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDefaultConfigs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDefaultConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
