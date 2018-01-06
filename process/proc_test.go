package process

import (
	"github.com/code-engine/go-engine/filesystem"
	"testing"
)

func TestRawList(t *testing.T) {
	tmpDir := filesystem.NewRelativeDir("./testdir")
	tmpDir.Create()
	defer tmpDir.Destroy()

	pd := ProcDir{Path: tmpDir.Path}

	tmpDir.NewFile("1", []byte{}, 0700)
	tmpDir.NewFile("namedwithstring", []byte{}, 0700)
	tmpDir.NewFile("1withstring", []byte{}, 0700)

	list, err := pd.RawList()

	if err != nil {
		t.Fatal(err)
	}

	listLength := len(list)

	if listLength != 1 {
		t.Fatalf("Incorrect number of results returned, expected 1, got %s", listLength)
	}

	if list[0] != 1 {
		t.Fatalf("Expected %d got %d", 1, list[0])
	}
}

func TestProcList(t *testing.T) {
	tmpDir := filesystem.NewRelativeDir("./testdir")
	tmpDir.Create()
	defer tmpDir.Destroy()

	pd := ProcDir{Path: tmpDir.Path}

	tmpDir.NewFile("2", []byte{}, 0700)
	tmpDir.NewFile("1", []byte{}, 0700)
	tmpDir.NewFile("123", []byte{}, 0700)
	tmpDir.NewFile("20", []byte{}, 0700)
	tmpDir.NewFile("10", []byte{}, 0700)

	list, err := pd.ProcList()

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
