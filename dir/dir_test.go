package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAttributeSetting(t *testing.T) {
	path := "/path/to/app_dir"

	dir := New(path)

	if dir.Path != path {
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

func TestJoin(t *testing.T) {
	dir := New("/path/to")

	expected := "/path/to/dir"
	actual := dir.Join("dir")

	if expected != actual {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestNewFile(t *testing.T) {
	testDir, err := filepath.Abs("./test_dir")

	if err != nil {
		t.Fatal(err)
	}

	dir := New(testDir)

	dir.Create()
	defer dir.Destroy()

	filename := "foo.txt"

	dir.NewFile(filename, []byte("Some text"), 0700)

	expectedPath := filepath.Join(testDir, filename)

	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Fatalf("File not found at %s", expectedPath)
	}
}

func TestExistsDirExists(t *testing.T) {
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

	if !dir.Exists() {
		t.Fatal("Expected true, got false")
	}
}

func TestExistsDirDoesNotExist(t *testing.T) {
	dir := NewRelative("./test_dir")

	if dir.Exists() {
		t.Fatal("Expected false, got true")
	}
}

func TestFileExistsFileDoesNotExist(t *testing.T) {
	dir := NewRelative("./test_dir")
	dir.Create()

	defer dir.Destroy()

	if dir.FileExists("testfile.txt") {
		t.Fatal("Expected false got true")
	}
}

func TestFileExistsFileExists(t *testing.T) {
	testDir, err := filepath.Abs("./test_dir")

	if err != nil {
		t.Fatal(err)
	}

	dir := New(testDir)

	dir.Create()

	defer dir.Destroy()

	filename := "testfile.txt"

	path := filepath.Join(testDir, filename)

	err = ioutil.WriteFile(path, []byte(""), 0700)

	if err != nil {
		t.Fatal(err)
	}

	if !dir.FileExists(filename) {
		t.Fatal("Expected true got false")
	}
}
