package clout

import (
	"io"
)

// Verbose is used to build messages.
//
// It's named after the klog equivalent for compatibility reasons.
type Verbose struct {
	verbosity MessageVerbosity
	printer   PrinterInterface
	enabled   bool
}

// Enabled returns true if the message will be printed.
// This can be used to avoid wasting CPU cycles on constructing messages that will never be displayed.
func (v *Verbose) Enabled() bool {
	return v.enabled
}

// Deprecationf prints a formatted Deprecation warning message.
func (v *Verbose) Deprecationf(format string, args ...interface{}) {
	if v.Enabled() {
		v.printer.Print(New(Deprecation, v.verbosity, format, args...))
	}
}

// Deprecationln prints a Deprecation warning message.
func (v *Verbose) Deprecationln(args ...interface{}) {
	if v.Enabled() {
		v.Deprecationf(argsToFormat(args), args...)
	}
}

// Deprecation is an alias for Deprecationln.
func (v *Verbose) Deprecation(args ...interface{}) {
	v.Deprecationln(args...)
}

// Warningf prints a formatted Warning message.
func (v *Verbose) Warningf(format string, args ...interface{}) {
	if v.Enabled() {
		v.printer.Print(New(Warning, v.verbosity, format, args...))
	}
}

// Warningln prints a Warning message.
func (v *Verbose) Warningln(args ...interface{}) {
	if v.Enabled() {
		v.Warningf(argsToFormat(args), args...)
	}
}

// Warning is an alias for Warningln.
func (v *Verbose) Warning(args ...interface{}) {
	v.Warningln(args...)
}

// Errorf prints a formatted Error message.
func (v *Verbose) Errorf(format string, args ...interface{}) {
	if v.Enabled() {
		v.printer.Print(New(Error, v.verbosity, format, args...))
	}
}

// Errorln prints an Error message.
func (v *Verbose) Errorln(args ...interface{}) {
	if v.Enabled() {
		v.Errorf(argsToFormat(args), args...)
	}
}

// Error is an alias for Errorln.
func (v *Verbose) Error(args ...interface{}) {
	v.Errorln(args...)
}

// Statusf prints a formatted Status message.
func (v *Verbose) Statusf(format string, args ...interface{}) {
	if v.Enabled() {
		v.printer.Print(New(Status, v.verbosity, format, args...))
	}
}

// Statusln prints a Status message.
func (v *Verbose) Statusln(args ...interface{}) {
	if v.Enabled() {
		v.Statusf(argsToFormat(args), args...)
	}
}

// Status is an alias for Statusln.
func (v *Verbose) Status(args ...interface{}) {
	v.Statusln(args...)
}

// Infof prints a formatted Info message.
func (v *Verbose) Infof(format string, args ...interface{}) {
	if v.Enabled() {
		v.printer.Print(New(Info, v.verbosity, format, args...))
	}
}

// Infoln prints an Info message.
func (v *Verbose) Infoln(args ...interface{}) {
	if v.Enabled() {
		v.Infof(argsToFormat(args), args...)
	}
}

// Info is an alias for Infoln.
func (v *Verbose) Info(args ...interface{}) {
	v.Infoln(args...)
}

// AsWriter creates an io.Writer that prints all incoming lines of text through the clout package.
// This is intended to convert the stdout and stderr of an executed command into Message objects.
//
// Example:
//
//     cmd := exec.Command("echo")
//     cmd.Stdout = clout.V(2).AsWriter(clout.Status)
//
func (v *Verbose) AsWriter(kind MessageKind) io.Writer {
	if !v.Enabled() {
		return io.Discard // If nothing will be printed anyways, just sinkhole incoming bytes.
	}

	return MessageWriter{
		Printer: v.printer,
		Converter: func(text string) *Message {
			msg := New(kind, v.verbosity, "%s", text)
			return &msg
		},
	}
}
