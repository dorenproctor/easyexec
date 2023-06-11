# easyexec

## Description

`easyexec` is a simple wrapper around the built-in `os/exec` library. While it's very flexible, I don't find the interface to `os/exec` to be friendly enough, so I made this small library to make running external commands in Go simple.

This library is built entirely using the core Go library. The only dependency is [github.com/stretchr/testify](https://github.com/stretchr/testify) for testing.

## Installing

`go get github.com/dorenproctor/easyexec`

## API

```go
// Wrapper for exec.Cmd that outputs Stdout, Stderr, and Error in a struct
func Run(command string, arg ...string) Output

// Wrapper for exec.Cmd that streams stdout, stderr, and error and outputs them in a struct
// Defaults are Stdout/Stderr pipes and not surrounding streamed stderr in red ANSI.
// If o is nil, a default CmdInput is used
func RunStream(o *Options, command string, arg ...string) Output

type Output struct {
    Err    error
    Stderr string
    Stdout string
}

// Options for RunStream()
type Options struct {
    // If nil, defaults to fmt.Print(s)
    PrintOut func(s string)
    // If nil, defaults to fmt.Fprint(os.Stderr, s)
    PrintErr func(s string)
    // If true, wrap text printed to StreamErr in red ANSI
    StreamErrRed bool
}
```


## Using

```go
package yours

import (
    "fmt"
    "strings"

    "github.com/dorenproctor/easyexec"
)

// Use the diff command built into the shell to compare 2 files or dirs
func Diff(src, dst string) error {
    o := easyexec.Run("diff", src, dst)
    s := o.Stdout + o.Stderr
    if s == "" {
        return nil
    }
    return fmt.Errorf(s)
}

func RunMyLongScript() error {
    // prints to the console in real time but also stores it in output
    o := easyexec.RunStream(nil, "python", "long_running_script.py")
    if o.Err != nil {
        return o.Err
    }
    // this (nonexistent) example script only ever prints to stdout
    // and sometimes returns successfully but prints FAIL in its output
    if strings.Contains(o.Stdout, "FAIL") {
        return fmt.Errorf("command failed: %s", o.Stdout)
    }
    return nil
}

func RunStreamWithStderrPrintedInRed() error {
    // prints to the console in real time but also stores it in output
    return easyexec.RunStream(
        &easyexec.Options{StreamErrRed: true},
        "bash",
        "my_script.sh",
        "--a-flag",
        "some",
        "args",
    ).Err
}

```