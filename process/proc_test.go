package process

import (
	"github.com/code-engine/go-engine/filesystem"
	"testing"
)

func TestList(t *testing.T) {
	tmpDir := filesystem.NewRelativeDir("./testdir")
	tmpDir.Create()
	defer tmpDir.Destroy()

	pd := ProcDir{Path: tmpDir.Path}

	tmpDir.NewFile("2", []byte{}, 0700)
	tmpDir.NewFile("1", []byte{}, 0700)
	tmpDir.NewFile("123", []byte{}, 0700)
	tmpDir.NewFile("20", []byte{}, 0700)
	tmpDir.NewFile("10", []byte{}, 0700)

	list, err := pd.List()

	if err != nil {
		t.Fatal(err)
	}

	listLength := len(list)

	if listLength != 5 {
		t.Fatalf("Incorrect number of results returned, expected 1, got %d", listLength)
	}

	expectedOrder := []int{1, 2, 10, 20, 123}

	for i, item := range list {
		if item != expectedOrder[i] {
			t.Fatalf("Expected %d, got %d", expectedOrder[i], item)
		}
	}
}

func TestExistsPIDExists(t *testing.T) {
	tmpDir := filesystem.NewRelativeDir("./testdir")
	tmpDir.Create()
	defer tmpDir.Destroy()

	pd := ProcDir{Path: tmpDir.Path}

	tmpDir.NewFile("1", []byte{}, 0700)

	exists, err := pd.Exists(1)

	if err != nil {
		t.Fatal(err)
	}

	if exists != true {
		t.Fatal("Expected true got false")
	}
}

func TestExistsPIDDoesNotExists(t *testing.T) {
	tmpDir := filesystem.NewRelativeDir("./testdir")
	tmpDir.Create()
	defer tmpDir.Destroy()

	pd := ProcDir{Path: tmpDir.Path}

	exists, err := pd.Exists(1)

	if err != nil {
		t.Fatal(err)
	}

	if exists != false {
		t.Fatal("Expected false got true")
	}
}
