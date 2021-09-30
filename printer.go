package clout

import (
	"fmt"
	"os"

	"go.eth-p.dev/clout/pkg/color"
)

// PrinterInterface processes and prints a Message.
// This can be used instead of Printer if more fine-grained control over messages are needed.
type PrinterInterface interface {
	Print(message Message)
}

// Printer is an implementation of PrinterInterface which prints to Output instances.
// Each MessageKind can be configured to use different Output instances.
type Printer struct {
	outputs  map[MessageKind]*Output
	fallback *Output
}

func (p *Printer) Print(message Message) {
	// Get the output for the message kind.
	output, ok := p.outputs[message.Kind()]
	if ok == false {
		output = p.fallback
	}

	// Write the message to the output.
	err := output.write(&message)
	if err != nil {
		panic(fmt.Errorf("failed to print message; err= %w", err))
	}
}

// SetOutput changes the default Output for all messages that are not handled by SetOutputForKind.
func (p *Printer) SetOutput(output Output) *Printer {
	p.fallback = &output
	return p
}

// SetOutputForKind changes the Output for all messages of a MessageKind.
func (p *Printer) SetOutputForKind(kind MessageKind, output Output) *Printer {
	p.outputs[kind] = &output
	return p
}

// NewPrinter creates a Printer with default settings.
// It will print all messages to stdout.
func NewPrinter() *Printer {
	stdout := OutputFromFile(os.Stdout)
	return (&Printer{outputs: make(map[MessageKind]*Output)}).
		SetOutput(stdout)
}

// NewPrinterWithDefaults creates a TerminalPrinter with default settings.
//
// It will print warnings and errors to stderr, and other messages to stdout.
// If stdout/stderr is a terminal, it will apply color output to those messages as well.
func NewPrinterWithDefaults(colors bool) *Printer {
	stderr := OutputFromFile(os.Stderr)

	return NewPrinter().
		SetOutputForKind(Warning, stderr.
			WithColor(optionallyColored(colors, color.Foreground(color.Yellow))).
			WithPrefix("warning:", optionallyColored(colors, color.Foreground(color.Yellow).Bold(true)))).
		SetOutputForKind(Deprecation, stderr.
			WithColor(optionallyColored(colors, color.Foreground(color.Yellow))).
			WithPrefix("deprecated:", optionallyColored(colors, color.Foreground(color.Yellow).Bold(true)))).
		SetOutputForKind(Error, stderr.
			WithColor(optionallyColored(colors, color.Foreground(color.Red))).
			WithPrefix("error:", optionallyColored(colors, color.Foreground(color.Red).Bold(true))))
}

func optionallyColored(enabled bool, c color.Style) color.Style {
	if !enabled {
		return color.Plain()
	}
	return c
}
