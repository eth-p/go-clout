package highlight

import (
	"go.eth-p.dev/clout/pkg/color"
)

// Highlight is an interface that can be used to change the style of formatted arguments.
type Highlight interface {

	// Value returns the value to be highlighted.
	Value() interface{}

	// Apply applies the highlighting style to a string.
	Apply(str string) string
}

// colorHighlight is an implementation of Highlight that uses color.Style to provide highlighting.
type colorHighlight struct {
	value interface{}
	style color.Style
}

func (c colorHighlight) Value() interface{} {
	return c.value
}

func (c colorHighlight) Apply(str string) string {
	return c.style.Apply(str)
}
