package clout

import (
	"fmt"

	"github.com/eth-p/clout/pkg/fitm"
	"github.com/eth-p/clout/pkg/highlight"
)

// formatText formats a Message's text into a string that can be printed.
// Colored text will be enabled or disabled based on the value of the colors parameter.
func formatText(message *Message, colors bool) string {
	var mitmFunc fitm.FormatMitm

	if colors {
		mitmFunc = fitmApplyColors
	} else {
		mitmFunc = fitmDiscardColors
	}

	return fitm.Sprintf(mitmFunc, message.Format(), message.FormatArgs()...)
}

// fitmApplyColors is a fitm.FormatMitm that applies colors from highlight.Highlight objects.
func fitmApplyColors(verb fitm.Verb, val interface{}) (fitm.Verb, interface{}) {
	if highlighter, ok := val.(highlight.Highlight); ok {
		formatted := fmt.Sprintf(verb.String(), highlighter.Value())
		highlighted := highlighter.Apply(formatted)

		return fitm.NewVerb("s"), highlighted
	}

	return verb, val
}

// fitmDiscardColors is a fitm.FormatMitm that discards colors from highlight.Highlight objects.
func fitmDiscardColors(verb fitm.Verb, val interface{}) (fitm.Verb, interface{}) {
	if highlighter, ok := val.(highlight.Highlight); ok {
		return verb, highlighter.Value()
	}

	return verb, val
}
