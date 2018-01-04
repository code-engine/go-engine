package filesystem

import (
	"path/filepath"
	"testing"
)

type Data struct {
	Name string
	Age  int
}

func TestFilePath(t *testing.T) {
	dirPath := "/path/to/dir"
	fileName := "test.yml"

	data := Data{}
	d := NewDir(dirPath)
	f := NewFile(fileName, data, &d)

	expectedPath := filepath.Join(dirPath, fileName)

	if f.Path() != expectedPath {
		t.Fatalf("Expected path to equal %s, got %s", expectedPath, f.Path())
	}
}

func TestFileCreate(t *testing.T) {
	d := NewRelativeDir("test_dir")
	f := NewFile("test.yml", Data{}, &d)

	d.Create()
	defer d.Destroy()

	err := f.Create()

	if err != nil {
		t.Fatal(err)
	}

	if !Exists(f.Path()) {
		t.Fatalf("File expected at %s, but file does not exist", f.Path())
	}
}
