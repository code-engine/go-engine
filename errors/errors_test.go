package utils

import (
	"errors"
	"fmt"
	"testing"

	"github.com/code-engine/go-engine/logger"
)

func TestCheckErrorWithError(t *testing.T) {
	defer logger.SetDefaultOutput()

	fakeOutput := logger.NewFakeOutput()

	logger.SetOutput(fakeOutput)

	errorMessage := "An error has occurred"

	err := errors.New(errorMessage)

	CheckError(err)

	if fakeOutput.Data.(error).Error() != errorMessage {
		t.Fatal("Actual and expected error messages do not match")
	}
}

func TestCheckErrorWithoutError(t *testing.T) {
	defer logger.SetDefaultOutput()

	fakeOutput := logger.NewFakeOutput()

	logger.SetOutput(fakeOutput)

	CheckError(nil)

	fmt.Println(fakeOutput.Data)

	if fakeOutput.Data != nil {
		t.Fatal("Error raised incorrectly")
	}
}
