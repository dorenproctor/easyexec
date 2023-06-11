package easyexec

import (
	"github.com/dorenproctor/easyexec/internal/helpers/runstream"
)

// Wrapper for exec.Cmd that streams stdout, stderr, and error and outputs them in a struct
//
// Defaults are Stdout/Stderr pipes and not surrounding streamed stderr in red ANSI.
//
// If o is nil, a default CmdInput is used
func RunStream(o *Options, command string, arg ...string) Output {
	options := runstream.Options(*o)
	return Output(runstream.RunStream(&options, command, arg...))
}
