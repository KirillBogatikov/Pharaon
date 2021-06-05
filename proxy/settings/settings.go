package settings

import (
	_ "embed"
	"gopkg.in/yaml.v2"
)

//go:embed settings.yml
var settings string

func Load() (Settings, error) {
	set := &YamlSettings{}
	err := yaml.Unmarshal([]byte(settings), set)
	if err != nil {
		return nil, err
	}

	result := make(Settings)
	for _, s := range set.Services {
		result[s.Name] = s
	}

	return result, nil
}
