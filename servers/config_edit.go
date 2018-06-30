// @CR: file ought to be named config_editor

package servers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"reflect"
	"github.com/alohen/dynamic_configuration_manager/config_handeling"
	"github.com/alohen/dynamic_configuration_manager/example_config/structs"
)

const (
	EditingUrlPrefix = "/edit/"
	ConfigEditError  = "Error editing example_config"
)

type ConfigEditingServer struct {
	configLoader *config_handeling.ConfigLoader
}

func NewConfigEditingServer(configLoader *config_handeling.ConfigLoader) http.Handler {
	return &ConfigEditingServer{
		configLoader: configLoader,
	}
}

type ConfigEditCommand struct {
	FileName       string
	OriginalConfig interface{}
	EditedConfig   interface{}
}

func (server *ConfigEditingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parser := structs.GetParser(filePath)
	if parser == nil {
		http.Error(w, MissingConfigError, 404)
		fmt.Println("Fail 1")
		return
	}

	config, err := server.configLoader.LoadFile(filePath)
	if err != nil {
		http.Error(w, MissingConfigError, 404)
		fmt.Println("Fail 2")
		return
	}

	configEdit := parser.ParseConfig(body)
	editedConfig := mergeConfigs(config, configEdit)

	data, err := json.Marshal(editedConfig)
	if err != nil {
		http.Error(w, ConfigEditError, 500)
		return
	}

	filePath = path.Join(server.configLoader.WorkingDirectory, config_handeling.ConfigPath, filePath)
	ioutil.WriteFile(filePath, data, os.ModePerm)
}

func mergeConfigs(originalConfig, EditedConfig interface{}) interface{} {
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
