package config

import (
	"github.com/jinzhu/configor"
	"log"
)

type Config struct {
	Router Router `yaml:"router"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Auth   Auth   `yaml:"auth"`
}

type Router struct {
	Port string `yaml:"port"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Passwd   string `yaml:"passwd"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Auth struct {
	Salt    string `yaml:"salt"`
	SignKey string `yaml:"signkey"`
}

var C *Config = &Config{}

func InitCfg(cfgFile string) {
	if cfgFile != "" {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C, cfgFile); err != nil {
			log.Fatal("init config fail", err)
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C); err != nil {
			log.Fatal("init config fail", err)
		}
	}
}
