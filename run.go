package easyexec

import (
	"github.com/dorenproctor/easyexec/internal/helpers/run"
)

// Wrapper for exec.Cmd that outputs stdout, stderr, and error in a struct
func Run(command string, arg ...string) Output {
	return Output(run.Run(command, arg...))
}
