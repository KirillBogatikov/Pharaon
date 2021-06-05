package config

import (
	"fmt"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
)

type DatabaseConfig struct {
	Host           string `yaml:"host"`
	Name           string `yaml:"name"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	MaxConnections int    `yaml:"max_connections"`
}

func (d *DatabaseConfig) PsqlConfig() *psql.DatabaseConfig {
	return &psql.DatabaseConfig{
		URL:            fmt.Sprintf("postgres://%s/%s?user=%s&password=%s", d.Host, d.Name, d.User, d.Password),
		MaxConnections: d.MaxConnections,
	}
}

type HttpConfig struct {
	BindAddress string `yaml:"bind_address"`
	ApiUrl      string `yaml:"api_url"`
}

type ServiceConfig struct {
	Name       string          `yaml:"name"`
	Database   *DatabaseConfig `yaml:"db"`
	HttpConfig *HttpConfig     `yaml:"http"`
}

type GlobalConfig struct {
	Services []ServiceConfig `yaml:"services"`
}
