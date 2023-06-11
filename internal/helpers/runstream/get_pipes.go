package runstream

import (
	"io"
	"os/exec"
)

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
