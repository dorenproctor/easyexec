package run

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/dorenproctor/easyexec/cmd/utils"
)

// Wrapper for exec.Cmd that outputs stdout, stderr, and error in a struct
func Run(command string, arg ...string) Output {
	if !utils.CommandExists(command) {
		return Output{Err: fmt.Errorf("command '%s' does not exist", command)}
	}
	cmd := exec.Command(command, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return Output{
		Err:    err,
		Stderr: stderr.String(),
		Stdout: stdout.String(),
	}
}

type Output struct {
	Err    error
	Stderr string
	Stdout string
}
