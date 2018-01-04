package filesystem

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func NewYAMLFile(name string, data interface{}, parentDir *Dir) YAMLFile {
	yamlFile := YAMLFile{}

	yamlFile.File.Name = name
	yamlFile.File.Data = data
	yamlFile.File.Dir = parentDir
	yamlFile.File.Perm = 0700

	return yamlFile
}

type YAMLFile struct {
	File
}

func (y YAMLFile) Create() error {
	if Exists(y.Path()) {
		message := fmt.Sprintf("File %s already exists", y.Path())

		return errors.New(message)
	}

	marshalled, err := yaml.Marshal(y.Data)

	if err != nil {
		return err
	}

	yamlOut := append([]byte("---\n"), marshalled...)

	err = ioutil.WriteFile(y.Path(), yamlOut, os.FileMode(y.Perm))

	if err != nil {
		return err
	}

	return nil
}
