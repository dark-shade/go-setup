package utils

import (
	"errors"
	"io/fs"
	"os"
)

// Exists returns whether the given file or directory exists, returns true if location exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateFile checks if the file already exists, if not then will create it and set it up
func CreateFile(name string, data []byte, perm fs.FileMode) error {
	_, err := os.Stat(name)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(name, data, perm)
		if err != nil {
			return err
		}
	} else {
		return errors.New(name + ": file already exists")
	}

	return nil
}
