package servers

import (
	"github.com/alohen/dynamic_configuration_manager/config_handeling"
	"net/http"
	"github.com/alohen/dynamic_configuration_manager/structs"
	"reflect"
	"strings"
	"fmt"
)

const(
	ReadConfigPrefix = "/read/"
	MissingConfigError = "No such config"
	PageBuildingError = "Couldn't build page"
)
type ConfigRetrieveServer struct {
	ConfigLoader *config_handeling.ConfigLoader
}

func (server *ConfigRetrieveServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, ReadConfigPrefix)
	config, err := server.ConfigLoader.LoadFile(filePath)
	if err != nil {
		http.Error(w, MissingConfigError,404)
		fmt.Println(err)
		return
	}

	p := createPage(config)
	page, err := p.Serialize()
	if err != nil {
		http.Error(w, PageBuildingError, 500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(*page))
}

func createPage(object interface{}) *structs.Page {
	objectValue := reflect.ValueOf(object).Elem()
	objectType := reflect.Indirect(reflect.ValueOf(object)).Type()
	pageFields := structs.Fields{}

	for index := 0; objectType.NumField() > index; index++ {
		structField := objectType.Field(index)
		pageField := structs.NewField(
			structField.Name,
			getInputType(structField.Type),
			objectValue.Field(index).Interface())

		pageFields = append(pageFields, pageField)
	}

	page := structs.NewPage(objectType.Name(), objectType.Name(), "return sendForm()", pageFields)
	return page
}

func getInputType(fieldType reflect.Type) string {
	var inputType string

	switch fieldType.Kind() {
	case reflect.Int:
		inputType = "number"
	case reflect.Uint:
		inputType = "number"
	case reflect.Float32:
		inputType = "number"
	case reflect.Float64:
		inputType = "number"
	default:
		inputType = "text"
	}

	return inputType
}
