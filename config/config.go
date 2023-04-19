package config

import (
	"io/ioutil"
  "log"
	"fmt"
  "gopkg.in/yaml.v3"
)

const CONFIG_FILE_PATH = "./config/config.yml"

type Config struct {
	Server struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Kafka struct {
		Topic struct {
			UpdateUser string `yaml:"updateUser"`
		} `yaml:"topic"`
		Brokers []string `yaml:"brokers"`
		GroupId string `yaml:"groupId"`
	} `yaml:"kafka"`
}

var cfg *Config

func NewConfig() (*Config, error) {
	buf, err := ioutil.ReadFile(CONFIG_FILE_PATH)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
			return nil, fmt.Errorf("in file %q: %v", CONFIG_FILE_PATH, err)
	}

	return c, nil
}

func init() {
	c, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg = c
}

func ServerName() string {
	return cfg.Server.Name
}

func ServerHost() string {
	return cfg.Server.Host
}

func ServerPort() string {
	return cfg.Server.Port
}

func KafkaTopicUpdateUser() string {
	return cfg.Kafka.Topic.UpdateUser
}

func KafkaBrokers() []string {
	return cfg.Kafka.Brokers
}

func KafkaGroupId() string {
	return cfg.Kafka.GroupId
}