package clout

import "strings"

// MessageWriter is an implementation of io.Writer that generates Message instances for each line of text received.
// Each generated Message will be sent directly to the PrinterInterface for printing.
type MessageWriter struct {
	Converter MessageConverter
	Printer   PrinterInterface
}

// MessageConverter converts a string of text into a Message.
// This is used by the MessageWriter when reading data.
type MessageConverter func(text string) *Message

func (w MessageWriter) Write(p []byte) (n int, err error) {
	text := strings.TrimRight(string(p), "\r\n")
	msg := w.Converter(text)

	if msg != nil {
		w.Printer.Print(*msg)
	}

	return len(p), nil
}
