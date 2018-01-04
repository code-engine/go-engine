package filesystem

import (
	"io/ioutil"
	"testing"
)

type Data struct {
	Name string
	Age  int
}

func TestYAMLFileCreate(t *testing.T) {
	data := Data{
		Name: "test",
		Age:  1,
	}

	d := NewRelativeDir("test_dir")
	f := NewYAMLFile("test.yml", data, &d)

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

	expectedFileContent := "---\nname: test\nage: 1\n"

	if fileContentString != expectedFileContent {
		t.Fatalf("Expected '%s', got '%s'", expectedFileContent, fileContent)
	}
}
