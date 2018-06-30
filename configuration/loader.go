package config_handeling

import (
	"github.com/alohen/dynamic_configuration_manager/example_config/structs"
	"io/ioutil"
	"path"
	"fmt"
)
const(
	ConfigPath = "example_config\\configuration"
)

type ConfigLoader struct {
	WorkingDirectory string
}

func (loader *ConfigLoader) LoadFile(filePath string) (interface{}, error ) {
	var parsedConfig interface{} = nil
	data, err := ioutil.ReadFile(path.Join(loader.WorkingDirectory,ConfigPath,filePath))
	if err != nil {
		return nil, err
	}

	parser := structs.GetParser(filePath)
	if parser == nil {
		return nil, fmt.Errorf("No Parser found for file: %v ", filePath)
	}

	parsedConfig = parser.ParseConfig(data)

	return parsedConfig, nil
}

