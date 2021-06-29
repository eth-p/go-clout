package color

// Color is an abstract representation of a terminal color code.
type Color int

const (
	None    Color = iota
	White   Color = iota
	Red     Color = iota
	Green   Color = iota
	Yellow  Color = iota
	Blue    Color = iota
	Magenta Color = iota
	Cyan    Color = iota
)

// Style is a struct of terminal text style attributes.
type Style struct {
	foreground Color
	background Color
	bold       bool
}

// Foreground creates a new Style with a foreground Color.
func Foreground(color Color) Style {
	return Plain().Foreground(color)
}

// Background creates a new Style with a background Color.
func Background(color Color) Style {
	return Plain().Background(color)
}

// Plain creates a new empty Style.
func Plain() Style {
	return Style{
		foreground: None,
		background: None,
		bold:       false,
	}
}

// Foreground applies a foreground color to the Style.
func (s Style) Foreground(color Color) Style {
	s.foreground = color
	return s
}

// Background applies a background color to the Style.
func (s Style) Background(color Color) Style {
	s.background = color
	return s
}

// Bold applies a bold attribute to the Style.
func (s Style) Bold(bold bool) Style {
	s.bold = bold
	return s
}
