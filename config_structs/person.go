package config_structs

import (
	"encoding/json"
)

func init() {
	StructToConfig["persons"] = &PersonConfig{}
}

var (
	StructToConfig = map[string]ConfigParser{}
)

type PersonConfig struct {
	Name    string
	SurName string
	Age     int
}

func (personConfig *PersonConfig) ParseConfig(config []byte) interface{} {
	var parsedConfig PersonConfig
	json.Unmarshal(config, &parsedConfig)
	return &parsedConfig
}
