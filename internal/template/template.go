package template

import (
	"embed"
	"os"
	_template "text/template"
)

//go:embed tmpl
var fs embed.FS

func Execute(data interface{}, templateFile string, fileString string, file *os.File) error {
	tmplHandler, err := _template.New(templateFile).Parse(fileString)
	if err != nil {
		return err
	}
	err = tmplHandler.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}

func GetFileString(templatePath string) (string, error) {
	fileString, err := fs.ReadFile(templatePath)
	if err != nil {
		return "", err
	}
	return string(fileString), nil
}
