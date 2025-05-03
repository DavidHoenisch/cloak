package utils

import (
	"cloak/internal/config"
)

func GetVarGroups(configPath string) ([]string, error) {

	var groups []string

	grps, err := config.ParseInConfigFile(configPath)
	if err != nil {
		return nil, err
	}

	var group = grps.Groups

	for _, v := range group {
		groups = append(groups, v.Name)
	}

	return groups, nil
}
