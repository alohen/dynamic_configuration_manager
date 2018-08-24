package editor

import (
	"encoding/json"
	"fmt"
	"github.com/alohen/dynamic_configuration_manager/configuration"
	"github.com/alohen/dynamic_configuration_manager/example_config/structs"
	"io/ioutil"
	"os"
	"path"
	"reflect"
)

type ConfigEditor struct {
	configLoader *configuration.ConfigLoader
}

func NewConfigEditor(configLoader *configuration.ConfigLoader) *ConfigEditor {
	return &ConfigEditor{
		configLoader: configLoader,
	}
}

func (editor *ConfigEditor) EditConfiguration(configPath string, configEdit []byte) error {
	originalConfig, err := editor.configLoader.LoadFile(configPath)
	if err != nil {
		return err
	}

	parsedConfigEdit, err := editor.parseEditedConfig(configPath, configEdit)
	if err != nil {
		return err
	}

	editedConfig := editor.mergeConfigs(originalConfig, parsedConfigEdit)
	err = editor.writeConfig(configPath, editedConfig)
	if err != nil {
		return err
	}

	return nil
}

func (editor *ConfigEditor) parseEditedConfig(configPath string, configEdit []byte) (interface{}, error) {
	parser := structs.GetParser(configPath)
	if parser == nil {
		return nil, configuration.NewParsingError(fmt.Errorf("parser not found for config file: %v", configPath))
	}

	editedConfig := parser.ParseConfig(configEdit)
	return editedConfig, nil
}

func (editor *ConfigEditor) mergeConfigs(originalConfig, EditedConfig interface{}) interface{} {
	config := reflect.ValueOf(originalConfig).Elem()
	editFields := reflect.ValueOf(EditedConfig).Elem()

	for i := 0; i < editFields.NumField(); i++ {
		originalField := config.Field(i)
		editField := editFields.Field(i)
		if editFields.Field(i).Interface() != reflect.Zero(editFields.Field(i).Type()).Interface() {
			originalField.Set(reflect.ValueOf(editField.Interface()))
		}
	}

	return config.Interface()
}

func (editor *ConfigEditor) writeConfig(configPath string, config interface{}) error {
	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}

	filePath := path.Join(editor.configLoader.ConfigPath, configPath)
	ioutil.WriteFile(filePath, data, os.ModePerm)

	return nil
}
