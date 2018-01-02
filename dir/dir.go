package dir

import (
	"errors"
	"fmt"
	"os"
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

func (a Dir) Create() error {
	if a.exists() {
		message := fmt.Sprintf("Directory %s already exists", a.path)
		return errors.New(message)
	}

	return os.MkdirAll(a.path, os.FileMode(a.perm))
}

func (a *Dir) SetPerm(perm int) {
	a.perm = perm
}

func (a Dir) Perm() int {
	return a.perm
}

func (a Dir) Path() string {
	return a.path
}

func (d Dir) Destroy() error {
	return os.RemoveAll(d.path)
}

func (a Dir) exists() bool {
	if _, err := os.Stat(a.path); err == nil {
		return true
	}

	return false
}
