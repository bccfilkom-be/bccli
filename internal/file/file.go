package file

import (
	"errors"
	"os"
	"strings"
)

func Create(name string) (*os.File, error) {
	if Exist(name) {
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

func Exist(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
