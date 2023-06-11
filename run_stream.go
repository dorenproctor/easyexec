package easyexec

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Wrapper for exec.Cmd that streams stdout, stderr, and error and outputs them in a struct
//
// Defaults are Stdout/Stderr pipes and not surrounding streamed stderr in red ANSI.
//
// If o is nil, a default CmdInput is used
func RunStream(o *Options, command string, arg ...string) Output {
	if !CommandExists(command) {
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

func getPipes(cmd *exec.Cmd) (stdout, stderr io.ReadCloser, err error) {
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err = cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	return stdout, stderr, nil
}

func stream(
	// text is read from here
	rc io.ReadCloser,
	// text is stored here
	output *string,
	// If true, wrap text printed to StreamErr in red ANSI
	makeRed bool,
	// Called to stream text as it comes. Must not be nil
	print func(s string),
) {
	r := bufio.NewReader(rc)
	line, err := r.ReadString('\n')
	for err == nil {
		s := line
		if makeRed {
			s = Red(line)
		}
		print(s)
		*output += line
		line, err = r.ReadString('\n')
	}
}

// Sets Options, PrintOut and PrintErr if they're nil.
// *Options must not be nil
func addDefaultFuncs(o *Options) {
	if o == nil {
		o = &Options{}
	}
	if o.PrintOut == nil {
		o.PrintOut = func(s string) {
			fmt.Print(s)
		}
	}
	if o.PrintErr == nil {
		o.PrintErr = func(s string) {
			fmt.Fprint(os.Stderr, s)
		}
	}
}
