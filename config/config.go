package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	. "github.com/code-engine/go-engine/errors"
)

type Config struct {
	path string
	perm int
}

func New(path string) Config {
	return Config{
		path: path,
		perm: 0700,
	}
}

func (c Config) Parse(out interface{}) {
	rawData, err := c.readRaw(c.path)

	CheckError(err)

	err = yaml.Unmarshal(rawData, out)

	CheckError(err)
}

func (c Config) Write(in interface{}) error {
	data, err := yaml.Marshal(in)

	ioutil.WriteFile(c.path, data, os.FileMode(c.perm))

	if err != nil {
		return err
	}

	return nil
}

func (c Config) readRaw(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
