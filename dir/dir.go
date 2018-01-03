package dir

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/code-engine/go-utils/errors"
)

func New(path string) Dir {
	return Dir{
		Path: path,
		Perm: 0700,
	}
}

func NewRelative(relativePath string) Dir {
	absolutePath, err := filepath.Abs(relativePath)

	CheckError(err)

	return Dir{
		Path: absolutePath,
		Perm: 0700,
	}
}

type Dir struct {
	Path    string
	Perm    int
	subdirs []Dir
}

func (d Dir) Create() error {
	if d.Exists() {
		message := fmt.Sprintf("Directory %s already exists", d.Path)
		return errors.New(message)
	}

	return os.MkdirAll(d.Path, os.FileMode(d.Perm))
}

func (d Dir) Destroy() error {
	if d.Path == "" {
		return errors.New("No path set")
	}

	d.DestroyFiles()
	return os.Remove(d.Path)
}

func (d Dir) DestroyFiles() {
	files, err := ioutil.ReadDir(d.Path)

	CheckError(err)

	for _, file := range files {
		if file.IsDir() {
			return
		}

		path := d.Join(file.Name())
		err = os.Remove(path)
		CheckError(err)
	}
}

func (d Dir) NewFile(filename string, data []byte, perm int) error {
	path := d.Join(filename)
	return ioutil.WriteFile(path, data, os.FileMode(perm))
}

func (d Dir) Join(path string) string {
	return filepath.Join(d.Path, path)
}

func (d Dir) Exists() bool {
	if _, err := os.Stat(d.Path); err == nil {
		return true
	}

	return false
}

func (d Dir) FileExists(filename string) bool {
	path := d.Join(filename)

	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}
