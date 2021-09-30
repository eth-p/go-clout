package main

import (
	"go.eth-p.dev/clout"
	"go.eth-p.dev/clout/pkg/highlight"
	"os"
)

func main() {

	// clout can apply colors to individual parameters passed to the `Verbose.*f` functions.
	// All you need to do is wrap the parameter in a highlight.Highlight object.
	clout.V(2).Infof("Hi, I'm the %s executable!", highlight.Magenta(os.Args[0]))

	// You can make things even more consistent by creating helper functions to automatically
	// wrap parameters in a Highlight object. Outside of this example, you would typically want to make these
	// functions in your module:
	path := func(path string) highlight.Highlight {
		return highlight.Cyan(path)
	}

	clout.V(2).Infof("Hello from %s!", path(os.Args[0]))

	// Highlighting also supports using non-string format verbs.
	// The Highlight wrapper is handled transparently when using the default implementation of the Printer.
	clout.V(2).Infof("The first 6 digits are Pi are: %1.5f", highlight.Green(3.1415926535))

	// Or if you want to use a custom Highlight implementation, that's possible too.
	// See custom_highlighter.go for an implementation example.
	clout.V(2).Infof("Custom highlighter? %t", MyHighlight{
		value:  true,
		prefix: "{",
		suffix: "}",
	})

}
