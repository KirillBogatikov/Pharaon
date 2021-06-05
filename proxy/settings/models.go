package settings

type Service struct {
	Name   string   `yaml:"name"`
	Opened []string `yaml:"opened"`
	Hidden []string `yaml:"hidden"`
}

type YamlSettings struct {
	Services []Service `yaml:"services"`
}

type Settings map[string]Service
