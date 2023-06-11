package runstream

import (
	"fmt"
	"os"
)

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
