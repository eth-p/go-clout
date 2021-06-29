// +build !windows

package color

import (
	"strings"
)

// Apply applies text styling to a string.
// For implementation purposes, this is using ANSI SGR sequences.
//
// https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
func (s Style) Apply(str string) string {
	var sb strings.Builder
	sb.Grow(11)

	fg := colorToAnsi[s.foreground]
	bg := colorToAnsi[s.background]

	if s.bold {
		appendAnsiParameter(&sb, "1")
	}

	if fg != "" {
		appendAnsiParameter(&sb, "3"+fg)
	}

	if bg != "" {
		appendAnsiParameter(&sb, "4"+bg)
	}

	// Return early if there's no s to be applied.
	if sb.Len() == 0 {
		return str
	}

	// Append the 'm' to the escape sequence and get it as a string.
	sb.WriteRune('m')
	ansi := sb.String()

	// Enable colors inside other colors by replacing the reset color with the parent color.
	// This only works for a depth of 1.
	str = strings.ReplaceAll(str, ansiReset, ansiReset+ansi)

	// Return a string with color codes surrounding it.
	return ansi + str + ansiReset
}

// appendAnsiParameter appends a parameter to a SGR escape sequence.
func appendAnsiParameter(builder *strings.Builder, parameter string) {
	if builder.Len() == 0 {
		builder.WriteString("\x1B[")
	} else {
		builder.WriteString(";")
	}
	builder.WriteString(parameter)
}

// ansiReset is the ANSI SGR escape sequence for resetting all colors back to default.
const ansiReset = "\x1B[0m"

// colorToAnsi is a lookup table that converts Color constants to ANSI SGR codes.
var colorToAnsi = [...]string{
	None:    "",
	White:   "9", // Technically default, but we avoid white and black because terminal BGs can be either.
	Red:     "1",
	Green:   "2",
	Yellow:  "3",
	Blue:    "4",
	Magenta: "5",
	Cyan:    "6",
}
