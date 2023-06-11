package runstream

import (
	"bufio"
	"io"

	"github.com/dorenproctor/easyexec/cmd/utils"
)

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
			s = utils.Red(line)
		}
		print(s)
		*output += line
		line, err = r.ReadString('\n')
	}
}
