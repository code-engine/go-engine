package utils

import (
	"github.com/code-engine/go-utils/logger"
)

func CheckError(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
