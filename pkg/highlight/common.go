package highlight

import (
	"go.eth-p.dev/clout/pkg/color"
)

// New creates a new Highlight from a color.Style
func New(value interface{}, style color.Style) Highlight {
	return colorHighlight{
		value: value,
		style: style,
	}
}

// Red highlights the value in red.
func Red(value interface{}) Highlight {
	return colorHighlight{
		value: value,
		style: color.Foreground(color.Red),
	}
}

// Green highlights the value in green.
func Green(value interface{}) Highlight {
	return colorHighlight{
		value: value,
		style: color.Foreground(color.Green),
	}
}

// Yellow highlights the value in yellow.
func Yellow(value interface{}) Highlight {
	return colorHighlight{
		value: value,
		style: color.Foreground(color.Yellow),
	}
}

// White highlights the value in white/black (depending on terminal background).
func White(value interface{}) Highlight {
	return colorHighlight{
		value: value,
		style: color.Foreground(color.White),
	}
}

// Cyan highlights the value in cyan.
func Cyan(value interface{}) Highlight {
	return colorHighlight{
		value: value,
		style: color.Foreground(color.Cyan),
	}
}

// Magenta highlights the value in magenta.
func Magenta(value interface{}) Highlight {
	return colorHighlight{
		value: value,
		style: color.Foreground(color.Magenta),
	}
}
