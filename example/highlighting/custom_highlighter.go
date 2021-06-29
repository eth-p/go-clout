package main

import "fmt"

// MyHighlight is a custom implementation of Highlight.
// If you want to use some other color package, this is how you would do it.
type MyHighlight struct {
	value  interface{}
	prefix string
	suffix string
}

// Value returns the value that will be highlighted.
// This will be run through fmt.Sprintf before being passed to Apply (if colors are enabled).
func (h MyHighlight) Value() interface{} {
	return h.value
}

// Apply gets called to apply the highlighting to the formatted value.
func (h MyHighlight) Apply(str string) string {
	return fmt.Sprintf("%s%s%s", h.prefix, str, h.suffix)
}
