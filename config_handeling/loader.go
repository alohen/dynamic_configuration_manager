package config_handeling

import (
	"fmt"
	"github.com/alohen/dynamic_configuration_manager/config_structs"
	"io/ioutil"
	"path"
)

const (
	ConfigPath = "config"
)

type ConfigLoader struct {
	ConfigDirectory string
}

func (loader *ConfigLoader) LoadFile(filePath string) (interface{}, error) {
	var parsedConfig interface{} = nil
	data, err := ioutil.ReadFile(path.Join(loader.ConfigDirectory, ConfigPath, filePath))
	if err != nil {
		return nil, err
	}

	parser := config_structs.GetParser(filePath)
	if parser == nil {
		return nil, fmt.Errorf("No Parser found for file: %v ", filePath)
	}

	parsedConfig = parser.ParseConfig(data)

	return parsedConfig, nil
}
