package template

import (
	"embed"
	"os"
	_template "text/template"
)

//go:embed tmpl
var tmpl embed.FS

func Execute(data interface{}, filename string, file *os.File) error {
	filename += ".tmpl"
	_content, err := content(filename)
	if err != nil {
		return err
	}
	tmpl, err := _template.New(filename).Parse(_content)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(file, data); err != nil {
		return err
	}
	return nil
}

func content(file string) (string, error) {
	fileString, err := tmpl.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(fileString), nil
}
