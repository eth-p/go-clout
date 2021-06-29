package clout

import (
	"io"
	"os"

	"github.com/eth-p/clout/pkg/color"
)

// Output is an io.Writer where formatted Messages are sent.
type Output struct {
	writer io.Writer
	colors bool

	terminator  string
	color       color.Style
	prefixColor color.Style
	prefix      string
}

// Clone creates a copy of the Output.
func (o Output) Clone() Output {
	return Output{
		writer:      o.writer,
		colors:      o.colors,
		color:       o.color,
		prefix:      o.prefix,
		prefixColor: o.prefixColor,
		terminator:  o.terminator,
	}
}

// WithColors creates a copy of the Output with colors enabled/disabled.
func (o Output) WithColors(colors bool) Output {
	clone := o.Clone()
	clone.colors = colors
	return clone
}

// WithColor creates a copy of the Output with a default text color.
// The default text color is applied to all messages that go through this output.
func (o Output) WithColor(color color.Style) Output {
	clone := o.Clone()
	clone.color = color
	return clone
}

// WithPrefix creates a copy of the Output with a prefix string.
func (o Output) WithPrefix(prefix string, color color.Style) Output {
	clone := o.Clone()
	clone.prefix = prefix
	clone.prefixColor = color
	return clone
}

// write writes a Message to the Output.
//
// This will convert the Message format string and arguments to a string,
// then format the whole message with a Formatter if one is provided.
func (o Output) write(message *Message) error {
	text := formatText(message, o.colors)
	prefix := o.prefix

	// Apply colors.
	if o.colors {
		text = o.color.Apply(text)
		prefix = o.prefixColor.Apply(prefix)
	}

	// Apply message prefix.
	if o.prefix != "" {
		text = prefix + " " + text
	}

	// Write to the output.
	_, err := o.writer.Write([]byte(text + o.terminator))
	return err
}

// OutputFromFile creates a Output from an os.File.
// If the file is a terminal, colors will be enabled.
func OutputFromFile(file *os.File) Output {
	colorsSupported := supportsColor(file)
	return OutputFromWriter(file).
		WithColors(colorsSupported)
}

// OutputFromWriter creates a Output from an io.Writer.
func OutputFromWriter(writer io.Writer) Output {
	return Output{
		writer:     writer,
		colors:     false,
		terminator: "\n",
	}
}
