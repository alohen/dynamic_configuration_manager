package structs

import (
	"bytes"
	"path"
	"text/template"
	"io/ioutil"
)

const(
	pageTemplate = "templates\\page.html"
)

type Page struct {
	Title  string
	Header string
	Action string
	Fields Fields
}

type serializeablePage struct {
	Title  string
	Header string
	Action string
	Fields string
}

func NewPage(title, header, action string, fields []*Field) *Page {
	return &Page{
		Title: title,
		Header: header,
		Action: action,
		Fields: fields,
	}
}

func (page *Page) Serialize() (*string, error) {
	serializedFields, err := page.Fields.Serialize()
	if err != nil {
		return nil, err
	}

	readyPage := serializeablePage{
		Title: page.Title,
		Header: page.Header,
		Action: page.Action,
		Fields: *serializedFields,
	}

	return readyPage.Serialize()
}

func(page *serializeablePage) Serialize() (*string, error) {
	var buffer bytes.Buffer

	text, err := ioutil.ReadFile(path.Join(WorkingDirectory,pageTemplate))
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("Page").Parse(string(text))
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&buffer,*page)
	if err != nil {
		return nil, err
	}

	htmlElement := buffer.String()
	return &htmlElement, nil
}
