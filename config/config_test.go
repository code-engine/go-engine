package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/code-engine/go-utils/dir"
)

type ParseTest struct {
	Name string
	Age  int
}

func createTestDir() (dir.Dir, error) {
	path, err := filepath.Abs("./testdir")

	if err != nil {
		return dir.Dir{}, err
	}

	tmpDir := dir.New(path)
	err = tmpDir.Create()

	if err != nil {
		return dir.Dir{}, err
	}

	return tmpDir, nil
}

func TestParse(t *testing.T) {
	tmpDir, err := createTestDir()

	if err != nil {
		t.Fatal(err)
	}

	defer tmpDir.Destroy()

	name := "foo"
	age := 10

	yamlString := fmt.Sprintf("---\nname: %s\nage: %d", name, age)

	filename := filepath.Join(tmpDir.Path(), "config.yml")
	fileContent := []byte(yamlString)

	ioutil.WriteFile(filename, fileContent, 0700)

	config := New(filename)
	out := ParseTest{}
	config.Parse(&out)

	if out.Name != name || out.Age != age {
		t.Fatal("Parsed object does not match expected")
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestWrite(t *testing.T) {
	tmpDir, err := createTestDir()

	if err != nil {
		t.Fatal(err)
	}

	defer tmpDir.Destroy()

	filename := filepath.Join(tmpDir.Path(), "config.yml")

	name := "bar"
	age := 22

	config := New(filename)
	testObject := ParseTest{
		Name: name,
		Age:  age,
	}

	config.Write(testObject)

	output, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Fatal(err)
	}

	expected := "name: bar\nage: 22\n"
	actual := string(output)

	if actual != expected {
		t.Log("Expected:", expected)
		t.Log("Got:", actual)
		t.Fatal("Output did not equal expected")
	}
}
