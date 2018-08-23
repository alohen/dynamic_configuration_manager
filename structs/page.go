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
	pageTemplate = "assets\\templates\\page.html"
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
		Title:  title,
		Header: header,
		Action: action,
		Fields: fields,
	}
}

func (page *Page) Serialize() ([]byte, error) {
	serializedFields, err := page.Fields.Serialize()
	if err != nil {
		return nil, err
	}

	readyPage := serializeablePage{
		Title:  page.Title,
		Header: page.Header,
		Action: page.Action,
		Fields: string(serializedFields),
	}

	return readyPage.Serialize()
}

func (page *serializeablePage) Serialize() ([]byte, error) {
	var buffer bytes.Buffer

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	text, err := ioutil.ReadFile(path.Join(cwd, pageTemplate))
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("Page").Parse(string(text))
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&buffer, *page)
	if err != nil {
		return nil, err
	}

	htmlElement := buffer.Bytes()
	return htmlElement, nil
}
