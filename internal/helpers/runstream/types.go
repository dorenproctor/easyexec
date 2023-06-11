package runstream

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
