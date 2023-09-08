package modules

import (
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Wifi string `yaml:"wifi"`
	Time int    `yaml:"time"`
	Url  string `yaml:"url"`
}

var ConfigData Config

func Init(data []byte) {
	if err := yaml.Unmarshal(data, &ConfigData); err != nil {
		log.Fatalf("无法解析 YAML 数据: %v", err)
	}
	log.Println(ConfigData)
}
