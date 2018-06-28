package structs

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

const (
	// @CR:	Create resource loader for dynamically building paths (also,
	//		use filepath.join for cross-platform-ness)
	fieldTemplate = "assets\\templates\\field.html"
)

type Field struct {
	Name       string
	InputType  string
	InputValue interface{}
}

type Fields []*Field

func (field *Field) Serialize() (*string, error) {
	var buffer bytes.Buffer

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	text, err := ioutil.ReadFile(path.Join(cwd, fieldTemplate))
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("Field").Parse(string(text))
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&buffer, *field)
	if err != nil {
		return nil, err
	}

	htmlElement := buffer.String()
	return &htmlElement, nil
}

func (fields Fields) Serialize() (*string, error) {
	var buffer bytes.Buffer
	for _, field := range fields {
		htmlField, err := field.Serialize()
		if err != nil {
			return nil, err
		}
		buffer.WriteString(*htmlField)
	}

	htmlFields := buffer.String()
	return &htmlFields, nil
}

func NewField(name, inputType string, inputValue interface{}) *Field {
	return &Field{
		Name:       name,
		InputType:  inputType,
		InputValue: inputValue,
	}
}
