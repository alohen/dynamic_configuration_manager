package page_builder

import (
	"github.com/alohen/dynamic_configuration_manager/configuration"
	"github.com/alohen/dynamic_configuration_manager/structs"
	"reflect"
)

type PageBuilder struct {
	configLoader *configuration.ConfigLoader
}

func (builder *PageBuilder) BuildEditingPage(configurationPath string) ([]byte, error) {
	config, err := builder.configLoader.LoadFile(configurationPath)
	if err != nil {
		return nil, err
	}

	p := createPage(config)
	page, err := p.Serialize()
	if err != nil {
		return nil, err
	}

	return page, nil
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

	//var fieldKindToInputType [Kind]string

	//fieldKindToInputType[reflect.Int] = "number"
	//fieldKindToInputType[reflect.Int] = "number"
	//fieldKindToInputType[reflect.Int] = "number"
	//fieldKindToInputType[reflect.Int] = "number"

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
