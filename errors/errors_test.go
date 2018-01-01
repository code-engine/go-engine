package utils

import (
	"errors"
	"testing"

	"github.com/code-engine/go-utils/logger"
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
	fakeOutput.Data = errors.New("Error that should not exist")

	logger.SetOutput(fakeOutput)

	CheckError(nil)

	if fakeOutput.Data != nil {
		t.Fatal("Error raised incorrectly")
	}
}
