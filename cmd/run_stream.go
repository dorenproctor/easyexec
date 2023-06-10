package cmdutils

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

// Wrapper for exec.Cmd that streams stdout, stderr, and error and outputs them in a struct
//
// Defaults are Stdout/Stderr pipes and not surrounding streamed stderr in red ANSI.
//
// If input is nil, a default CmdInput is used
func RunStream(input *Input, command string, arg ...string) Output {
	if !CommandExists(command) {
		return Output{Err: fmt.Errorf("command '%s' does not exist", command)}
	}
	cmd := exec.Command(command, arg...)

	if input == nil {
		input = &Input{}
	}
	if input.StreamOut == nil {
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return Output{Err: err}
		}
		input.StreamOut = stdout
	}
	if input.StreamErr == nil {
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return Output{Err: err}
		}
		input.StreamErr = stderr
	}

	output := Output{}
	go stream(input.StreamOut, &output.Stdout, false)
	go stream(input.StreamErr, &output.Stderr, input.StreamErrRed)
	if err := cmd.Start(); err != nil {
		return Output{Err: err}
	}
	output.Err = cmd.Wait()
	return output
}

func stream(rc io.ReadCloser, output *string, makeRed bool) {
	r := bufio.NewReader(rc)
	line, err := r.ReadString('\n')
	for err == nil {
		if makeRed {
			fmt.Print(Red(line))
		}
		*output += line
		line, err = r.ReadString('\n')
	}
}
