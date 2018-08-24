package configuration

import (
	"github.com/alohen/dynamic_configuration_manager/example_config/structs"
	"io/ioutil"
	"path"
	"fmt"
)

type ConfigLoader struct {
	WorkingDirectory string
	ConfigPath string
}

func NewConfigLoader(configPath string) *ConfigLoader {
	return &ConfigLoader{
		ConfigPath: configPath,
	}
}
func (loader *ConfigLoader) LoadFile(filePath string) (interface{}, error ) {
	var parsedConfig interface{} = nil
	data, err := ioutil.ReadFile(path.Join(loader.ConfigPath,filePath))
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

