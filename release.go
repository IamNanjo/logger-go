//go:build !debug

package logger

import "log"

type EmptyWriter struct{}

func (e EmptyWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

var Debug = log.New(new(EmptyWriter), "", 0)
