package util

import (
	"errors"
	"os"
	"strings"
)

func CreateFile(name string) error {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		if strings.Contains(name, "/") {
			path := name[:strings.LastIndex(name, "/")]

			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		}

		_, err = os.Create(name)
		if err != nil {
			return err
		}
	} else {
		return errors.New("File already exist")
	}

	return nil
}
