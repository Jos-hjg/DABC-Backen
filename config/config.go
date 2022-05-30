package config

import (
	"github.com/jinzhu/configor"
	"log"
)

type Config struct {
	Router Router
	Mysql  Mysql
}

type Router struct {
	Port string
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Passwd   string `yaml:"passwd"`
	Database string `yaml:"database"`
}

var C *Config = &Config{}

func Init(cfgFile string) {
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
