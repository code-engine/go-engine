package filesystem

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestFilePath(t *testing.T) {
	dirPath := "/path/to/dir"
	fileName := "test.txt"
	d := NewDir(dirPath)
	f := NewTextFile(fileName, "", &d)

	expectedPath := filepath.Join(dirPath, fileName)

	if f.Path() != expectedPath {
		t.Fatalf("Expected path to equal %s, got %s", expectedPath, f.Path())
	}
}

func TestFileCreate(t *testing.T) {
	c := "Content"
	d := NewRelativeDir("test_dir")
	f := NewTextFile("test.txt", c, &d)

	d.Create()
	defer d.Destroy()

	err := f.Create()

	if err != nil {
		t.Fatal(err)
	}

	if !Exists(f.Path()) {
		t.Fatalf("File expected at %s, but file does not exist", f.Path())
	}

	fileContent, err := ioutil.ReadFile(f.Path())

	if err != nil {
		t.Fatal(err)
	}

	fileContentString := string(fileContent)

	expectedFileContent := c

	if fileContentString != expectedFileContent {
		t.Fatalf("Expected '%s', got '%s'", expectedFileContent, fileContent)
	}
}
