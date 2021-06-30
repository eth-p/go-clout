package clout

import (
	"bytes"
	"io"
	"strings"
)

// messageWriter is an implementation of io.Writer that generates Message instances for each line of text received.
// Each generated Message will be sent directly to the PrinterInterface for printing.
type messageWriter struct {
	Converter MessageConverter
	Printer   PrinterInterface

	buffer bytes.Buffer
}

// MessageConverter converts a string of text into a Message.
// This is used by the messageWriter when reading data.
type MessageConverter func(text string) *Message

func (w *messageWriter) Write(p []byte) (n int, err error) {
	pStr := w.buffer.String() + string(p)
	split := strings.Split(pStr, "\n")

	messages := split[:len(split)-1]
	last := split[len(split)-1]

	// If the last split isn't empty, we didn't end on a newline.
	// Append the remaining string to the buffer.
	if last != "" {
		w.buffer.Reset()
		w.buffer.WriteString(last)
	}

	// Process each line as its own message.
	for _, text := range messages {
		text = strings.TrimRight(text, "\r\n")
		msg := w.Converter(text)
		if msg != nil {
			w.Printer.Print(*msg)
		}
	}

	return len(p), nil
}

// MessageWriter creates an io.Writer that generates and prints Message instances for each line of text received.
func MessageWriter(converter MessageConverter, printer PrinterInterface) io.Writer {
	return &messageWriter{
		Printer: printer,
		Converter: func(text string) *Message {
			message := converter(text)

			// Ensure it doesn't print anything above the configured verbosity.
			if message == nil || message.verbosity > GetVerbosity() {
				return nil
			}

			return message
		},
	}
}
