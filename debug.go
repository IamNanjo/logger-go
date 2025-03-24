//go:build !release

package logger

import (
	"log"
	"os"
)

// Writes to stdout if release build tag is not used
var Debug = log.New(os.Stdout, "DEBUG: ", logFlags)
