package config

import (
	embed "github.com/Projector-Solutions/Pharaon-config/embed"
	"gopkg.in/yaml.v2"
	"os"
)

var Global *GlobalConfig

func IsProd() bool {
	f := os.Getenv("ENV")
	return f == "PROD"
}

func Service(name string) *ServiceConfig {
	for _, s := range Global.Services {
		if s.Name == name {
			return &s
		}
	}

	return nil
}

func LoadConfig(dev, prod string, s interface{}) error {
	var bytes []byte
	if IsProd() {
		bytes = []byte(prod)
	} else {
		bytes = []byte(dev)
	}

	return yaml.Unmarshal(bytes, s)
}

func init() {
	Global = &GlobalConfig{}
	err := LoadConfig(embed.ConfigDev, embed.ConfigProd, Global)
	if err != nil {
		panic(err)
	}
}
