//go:build debug

package logger

import (
	"log"
	"os"
)

// Writes to stdout if debug build tag is used
var Debug = log.New(os.Stdout, "DEBUG: ", logFlags)
