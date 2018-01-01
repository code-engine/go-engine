package logger

import (
	"testing"
)

func TestPrinting(t *testing.T) {
	defer SetDefaultOutput()

	newOutput := NewFakeOutput()
	SetOutput(newOutput)

	expected := "Foobar"

	Print(expected)

	if newOutput.Data != expected {
		t.Fatal("Actual does not match expected")
	}
}
