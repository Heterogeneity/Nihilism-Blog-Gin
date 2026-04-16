package core

import (
	"gopkg.in/yaml.v3"
	"log"
	"server/config"
	"server/utils"
)

// InitConf 从 YAML 文件加载配置
func InitConf() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYAML()
	if err != nil {
		log.Fatalf("加载yaml文件失败：%v", err)
	}
	if err = yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("fail to unmarshal yaml conf: %v", err)
	}
	return c
}
