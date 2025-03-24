package logger

import (
	"log"
	"os"
)

var logFlags = log.Ldate | log.Ltime

// Writes to stdout
var Out = log.New(os.Stdout, "OUT: ", logFlags)

// Writes to stderr
var Err = log.New(os.Stderr, "ERR: ", logFlags)
