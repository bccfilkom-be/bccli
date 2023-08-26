package template

import "embed"

//go:embed file-template
var fs embed.FS

func GetFileString(templatePath string) (string, error) {

	fileString, err := fs.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	return string(fileString), nil

}
