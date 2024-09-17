package file

import (
	"errors"
	"os"
	"strings"
)

func Create(name string) (*os.File, error) {
	_, err := os.Stat(name)
	if os.IsExist(err) {
		return nil, errors.New("file already exist")
	}

	if strings.Contains(name, "/") {
		path := name[:strings.LastIndex(name, "/")]
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}
