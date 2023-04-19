package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type EnvConfig struct {
	JwtServiceSecret       string `envconfig:"JWT_SERVICE_SECRET"`
	MysqlAddr    string `envconfig:"MYSQL_ADDR"`
	MysqlDB     string `envconfig:"MYSQL_DATABASE"`
	MysqlUser     string `envconfig:"MYSQL_USER"`
	MysqlPassword     string `envconfig:"MYSQL_PASS"`
}

var envCfg *EnvConfig

func NewEnvConfig() (*EnvConfig, error) {
	var c EnvConfig
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func init() {
	c, err := NewEnvConfig()
	if err != nil {
		log.Fatal(err)
	}
	envCfg = c
}

func EnvJwtServiceSecrect() string {
	return envCfg.JwtServiceSecret
}

func EnvMysqlAddr() string {
	return envCfg.MysqlAddr
}

func EnvMysqlDB() string {
	return envCfg.MysqlDB
}

func EnvMysqlUser() string {
	return envCfg.MysqlUser
}

func EnvMysqlPassword() string {
	return envCfg.MysqlPassword
}
