package filesystem

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

// Error if parent does not exist

func NewFile(name string, data interface{}, parentDir *Dir) File {
	return File{
		Name: name,
		Data: data,
		Dir:  parentDir,
		Perm: 0700,
	}
}

type File struct {
	Data interface{}
	Name string
	Dir  *Dir
	Perm int
}

func (f File) Create() error {
	if Exists(f.Path()) {
		message := fmt.Sprintf("File %s already exists", f.Path())

		return errors.New(message)
	}

	dataString := f.Data.(string)

	err := ioutil.WriteFile(f.Path(), []byte(dataString), os.FileMode(f.Perm))

	if err != nil {
		return err
	}

	return nil
}

func (f File) Path() string {
	return filepath.Join(f.Dir.Path, f.Name)
}

func (f File) Exists() bool {
	return Exists(f.Path())
}
