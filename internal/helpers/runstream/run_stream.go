package runstream

import (
	"fmt"
	"os/exec"

	"github.com/dorenproctor/easyexec/cmd/utils"
)

// Wrapper for exec.Cmd that streams stdout, stderr, and error and outputs them in a struct
//
// Defaults are Stdout/Stderr pipes and not surrounding streamed stderr in red ANSI.
//
// If o is nil, a default CmdInput is used
func RunStream(o *Options, command string, arg ...string) Output {
	if !utils.CommandExists(command) {
		return Output{Err: fmt.Errorf("command '%s' does not exist", command)}
	}

	if o == nil {
		o = &Options{}
	}
	addDefaultFuncs(o)

	cmd := exec.Command(command, arg...)
	stdout, stderr, err := getPipes(cmd)
	if err != nil {
		return Output{Err: err}
	}
	output := Output{}
	go stream(stdout, &output.Stdout, false, o.PrintOut)
	go stream(stderr, &output.Stderr, o.StreamErrRed, o.PrintErr)

	if err := cmd.Start(); err != nil {
		return Output{Err: err}
	}
	output.Err = cmd.Wait()
	return output
}
