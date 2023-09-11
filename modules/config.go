package modules

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Wifi     string `yaml:"wifi"`
	Time     int    `yaml:"time"`
	BaseUrl  string `yaml:"url"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
	Url      string `yaml:"-"`
}

var ConfigData Config

func Init(data []byte) {
	if err := yaml.Unmarshal(data, &ConfigData); err != nil {
		log.Fatalf("无法解析 YAML 数据: %v", err)
	}
	ConfigData.Url = ConfigData.BaseUrl + fmt.Sprintf("?username=%s&password=%s", ConfigData.Account, ConfigData.Password)
	log.Println(ConfigData)
}
