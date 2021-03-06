package filesystem

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"sort"

	. "github.com/code-engine/go-engine/errors"
)

func SortDirectories(dirs []Dir) {
	sort.SliceStable(dirs, func(i, j int) bool {
		return dirs[i].Priority < dirs[j].Priority
	})
}

func NewDir(path string) Dir {
	return Dir{
		Path:     path,
		Perm:     0700,
		Priority: -1,
	}
}

func NewRelativeDir(relativePath string) Dir {
	absolutePath, err := filepath.Abs(relativePath)

	CheckError(err)

	return NewDir(absolutePath)
}

func NewHomeRelativeDir(dirname string) Dir {
	currentUser, err := user.Current()

	CheckError(err)

	currentUserHomeDir := currentUser.HomeDir
	path := filepath.Join(currentUserHomeDir, dirname)

	return NewDir(path)
}

type Dir struct {
	Path     string
	Perm     int
	Priority int
	Subdirs  []Dir
	Files    []File
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

	err := d.DestroyFiles()

	if err != nil {
		return err
	}

	return os.Remove(d.Path)
}

func (d Dir) DestroyFiles() error {
	files, err := ioutil.ReadDir(d.Path)

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := d.Join(file.Name())
		err = os.Remove(path)

		if err != nil {
			return err
		}
	}

	return nil
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
