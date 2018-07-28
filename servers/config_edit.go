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
	"strings"

	"github.com/alohen/dynamic_configuration_manager/config_handeling"
	"github.com/alohen/dynamic_configuration_manager/config_structs"
)

const (
	EditingUrlPrefix = "/edit/"
	ConfigEditError  = "Error editing config"
)

type ConfigEditingServer struct {
	ConfigLoader *config_handeling.ConfigLoader
}

type ConfigEditCommand struct {
	FileName       string
	OriginalConfig interface{}
	EditedConfig   interface{}
}

func (server *ConfigEditingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, EditingUrlPrefix)

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parser := config_structs.GetParser(filePath)
	if parser == nil {
		http.Error(w, MissingConfigError, 404)
		fmt.Println("Fail 1")
		return
	}

	config, err := server.ConfigLoader.LoadFile(filePath)
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

	filePath = path.Join(server.ConfigLoader.ConfigDirectory, config_handeling.ConfigPath, filePath)
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
