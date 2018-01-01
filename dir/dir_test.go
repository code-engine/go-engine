package dir

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAttributeSetting(t *testing.T) {
	path := "/path/to/app_dir"

	dir := New(path)

	if dir.Path() != path {
		t.Fatal("Paths do not match")
	}
}

func TestCreateDirDoesNotExist(t *testing.T) {
	testDir, err := filepath.Abs("./test_dir")

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(testDir)

	dir := New(testDir)
	dir.Create()

	fileStat, err := os.Stat(testDir)

	if os.IsNotExist(err) {
		t.Fatalf("Path %s does not exist", testDir)
	}

	if fileStat.Mode().Perm() != os.FileMode(0700) {
		t.Fatal("Filemode is incorrect")
	}
}

func TestCreateDirExists(t *testing.T) {
	testDir, err := filepath.Abs("./test_dir")

	if err != nil {
		t.Fatal(err)
	}

	err = os.MkdirAll(testDir, 0700)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(testDir)

	dir := New(testDir)
	err = dir.Create()

	if err == nil {
		t.Fatal("Expected Create() to return an error, no error returned.")
	}
}

func TestSetPerm(t *testing.T) {
	dir := New("/path/to/dir")
	perm := 0600
	dir.SetPerm(perm)

	if dir.Perm() != perm {
		t.Fatalf("Permission 0%o does not equal expected 0%o", dir.Perm(), perm)
	}
}
