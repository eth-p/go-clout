package main

import (
	"fmt"
	"go.eth-p.dev/clout"
	"go.eth-p.dev/clout/pkg/fitm"
	"go.eth-p.dev/clout/pkg/highlight"
)

// MyPrinter is a custom implementation of PrinterInterface.
// This is much more flexible than the default implementation, but it involves some boilerplate to work with.
type MyPrinter struct {
}

// Print is called when a message should be printed.
// Any messages of higher verbosity than the minimum will be filtered out before this is called.
func (m MyPrinter) Print(message clout.Message) {
	// Use fitm to apply the Highlight wrappers.
	// This does not consider whether or not the terminal supports colors, and applies them unconditionally.
	text := fitm.Sprintf(applyHighlights, message.Format(), message.FormatArgs()...)

	// Print the formatted text.
	fmt.Printf("[V(%d)] %s\n", message.Verbosity(), text)
}

// applyHighlights is a fitm.FormatMitm function which applies highlight.Highlight wrappers to format arguments.
// This is necessary to allow highlighting to be formatted transparently.
func applyHighlights(verb fitm.Verb, val interface{}) (fitm.Verb, interface{}) {
	wrapper, ok := val.(highlight.Highlight)

	// If it's not a Highlight, return the verb and value as-is.
	if !ok {
		return verb, val
	}

	// Format the wrapped value with the verb.
	formatted := verb.Format(wrapper.Value())

	// Apply the highlight to the formatted value.
	highlighted := wrapper.Apply(formatted)

	// Return highlighted value as a preformatted string.
	return fitm.Preformatted(highlighted)
}
