package config

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

// Init 解析app配置
func Init(container Conf, customerPath ...string) error {
	filePath := configFilePath
	if len(customerPath) > 0 {
		filePath = customerPath[0]
	}
	fd, err := os.Open(filePath)
	if err != nil {
		return err
	}
	err = InitFromYaml(fd, container)
	if err != nil {
		return err
	}

	if container.AppConfig() == nil {
		container.SetAppConfig(defaultAppConfig)
	}
	if container.LogConfig() == nil {
		container.SetLogConfig(defaultLogConfig)
	}
	return nil
}

var configFilePath = "./application.yaml"

// InitFromJson 从json中解析配置
func InitFromJson(reader io.Reader, container interface{}) error {
	content, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, container)
	if err != nil {
		return err
	}
	return nil
}

// InitFromYaml 从yaml中解析配置
func InitFromYaml(reader io.Reader, container interface{}) error {
	content, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, container)
	if err != nil {
		return err
	}
	return nil
}
