package dir

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func New(path string) Dir {
	return Dir{
		path: path,
		perm: 0700,
	}
}

type Dir struct {
	path    string
	perm    int
	subdirs []Dir
}

func (d Dir) Create() error {
	if d.exists() {
		message := fmt.Sprintf("Directory %s already exists", d.path)
		return errors.New(message)
	}

	return os.MkdirAll(d.path, os.FileMode(d.perm))
}

func (d *Dir) SetPerm(perm int) {
	d.perm = perm
}

func (d Dir) Perm() int {
	return d.perm
}

func (d Dir) Path() string {
	return d.path
}

func (d Dir) Destroy() error {
	return os.RemoveAll(d.path)
}

func (d Dir) NewFile(filename string, data []byte, perm int) error {
	path := d.Join(filename)
	return ioutil.WriteFile(path, data, os.FileMode(perm))
}

func (d Dir) Join(path string) string {
	return filepath.Join(d.path, path)
}

func (d Dir) exists() bool {
	if _, err := os.Stat(d.path); err == nil {
		return true
	}

	return false
}
