package strings

import (
	"fmt"
	"testing"
)

func TestMultiLineToSlice(t *testing.T) {
	str := "This is line 1\nThis is line 2\nThis is line 3"

	out := MultiLineToSlice(str)

	for i, item := range out {
		expected := fmt.Sprintf("This is line %d", i+1)
		if item != expected {
			t.Fatalf("Expected '%s' got '%s'", expected, item)
		}
	}
}

func TestExtract(t *testing.T) {
	str := "Column1 Column2  Column3   Column4    "
	rgx := `(\w+)\s+(\w+)\s+(\w+)\s+(\w+)\s+`

	out, err := Extract(rgx, str)

	if err != nil {
		t.Fatal(err)
	}

	for i, item := range out {
		expected := fmt.Sprintf("Column%d", i+1)

		if item != expected {
			t.Fatalf("Expected %s got %s", expected, item)
		}
	}
}
