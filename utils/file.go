package utils

import (
	"github.com/spf13/afero"
	"os"
	"path/filepath"
)

func WriteToFile(filePath string, data []byte) error {
	var AppFs = afero.NewOsFs()

	dir, _ := filepath.Split(filePath)
	exists, err := afero.DirExists(AppFs, dir)
	if err != nil {
		return err
	}
	if !exists {
		err := AppFs.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err = afero.WriteFile(AppFs, filePath, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
