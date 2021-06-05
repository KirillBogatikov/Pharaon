package proxy

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	"github.com/Projector-Solutions/Pharaon-config/embed"
)

type Config struct {
	MongoURL string `yaml:"mongo_url"`
}

var Service *config.ServiceConfig
var Http *config.HttpConfig
var Env *Config

func init() {
	Service = config.Service("proxy")
	if Service == nil {
		panic("proxy config not found!")
	}

	Http = Service.HttpConfig
	if Http == nil {
		panic("proxy http config required!")
	}

	Env = &Config{}
	err := config.LoadConfig(embed.ProxyDev, embed.ProxyProd, Env)
	if err != nil {
		panic(err)
	}
}
