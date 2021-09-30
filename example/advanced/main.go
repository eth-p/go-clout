package main

import (
	"go.eth-p.dev/clout"
	"go.eth-p.dev/clout/pkg/color"
	"os"
)

// clout is opinionated about it's defaults, but it's not uncompromising.
// If you don't like any of the default settings, you can change them.

func main() {

	// If you don't like the prefixes in the default printer, you can configure a different printer:
	stdout := clout.OutputFromFile(os.Stdout)
	printer := clout.NewPrinter().
		SetOutputForKind(clout.Status, stdout.
			WithPrefix("*", color.Background(color.Magenta)).
			WithColor(color.Foreground(color.Magenta)))

	clout.SetPrinter(printer)
	clout.V(2).Statusf("status message")

	// Or if you don't like the printer at all, you can make your own printer.
	// See custom_printer.go for an implementation example.
	clout.SetPrinter(MyPrinter{})
	clout.V(2).Warningf("warning message")

}
