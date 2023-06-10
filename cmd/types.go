package cmdutils

import "io"

type Output struct {
	Err    error
	Stderr string
	Stdout string
}

type Input struct {
	// Defaults to Stdout pipe
	StreamOut io.ReadCloser
	// Defaults to Stderr pipe
	StreamErr io.ReadCloser
	// If true, wrap text printed to StreamErr in red ANSI
	StreamErrRed bool
}
