package clout

import "fmt"

// MessageKind represents the kind of message.
type MessageKind int

const (
	// Status represents an update to the program's current status.
	// This should be used to inform the user that the program is performing a new procedure.
	Status MessageKind = iota

	// Info represents an informational message about something.
	// This should be used to let the user know about the state of an object.
	Info MessageKind = iota

	// Warning represents a warning about a potential issue.
	// This should be used to warn the user about a minor problem.
	Warning MessageKind = iota

	// Deprecation represents a warning about a feature which will be removed or unsupported in the future.
	// This should be used when the user needs to be told that they're relying on a deprecated feature.
	Deprecation MessageKind = iota

	// Error represents a severe error.
	// This should be used to tell the user that the program was unable to complete an action.
	Error MessageKind = iota

	// Custom is a custom message kind.
	// This can be used with a custom printer to format special messages.
	Custom MessageKind = iota
)

// MessageVerbosity represents the verbosity level of a Message.
//
// This follows the Kubernetes log level convention at
// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/logging.md:
//
//   V(0) - Programmer errors, logging extra info about a panic, cli argument handling
//   V(1) - Information about config, errors
//   V(2) - System state, log messages
//   V(3) - Extended info about system state changes
//   V(4) - Logging in "thorny parts of code"
//   V(5) - Trace level verbosity
type MessageVerbosity int

// Message is a structured representation of a printable message.
type Message struct {
	format     string
	formatArgs []interface{}
	verbosity  MessageVerbosity
	kind       MessageKind
}

// String formats the message and returns its string.
func (m Message) String() string {
	return fmt.Sprintf(m.format, m.formatArgs...)
}

// Format returns the message's formatting string.
func (m Message) Format() string {
	return m.format
}

// FormatArgs returns the message's formatting args.
func (m Message) FormatArgs() []interface{} {
	return m.formatArgs
}

// Verbosity returns the message verbosity.
func (m Message) Verbosity() MessageVerbosity {
	return m.verbosity
}

// Kind returns the message kind.
func (m Message) Kind() MessageKind {
	return m.kind
}

// New creates a new Message.
func New(kind MessageKind, verbosity MessageVerbosity, format string, args ...interface{}) Message {
	return Message{
		format:     format,
		formatArgs: args,
		verbosity:  verbosity,
		kind:       kind,
	}
}

// defaultVerbosity is the lowest MessageVerbosity that will be shown by default.
const defaultVerbosity MessageVerbosity = 2
