package clout

import (
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testPrinter struct {
	messages []Message
}

func (p *testPrinter) Print(message Message) {
	p.messages = append(p.messages, message)
}

func TestVerbose(t *testing.T) {
	tests := map[string]struct {
		expected  []Message
		verbosity MessageVerbosity
		fn        func(v Verbose)
	}{
		"Infof": {
			expected: []Message{{
				format:     "hello %s",
				formatArgs: []interface{}{"infof"},
				kind:       Info,
			}},
			fn: func(v Verbose) {
				v.Infof("hello %s", "infof")
			},
		},
		"Infoln": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"infoln"},
				kind:       Info,
			}},
			fn: func(v Verbose) {
				v.Infoln("infoln")
			},
		},
		"Info": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"info"},
				kind:       Info,
			}},
			fn: func(v Verbose) {
				v.Info("info")
			},
		},
		"Statusf": {
			expected: []Message{{
				format:     "hello %s",
				formatArgs: []interface{}{"statusf"},
				kind:       Status,
			}},
			fn: func(v Verbose) {
				v.Statusf("hello %s", "statusf")
			},
		},
		"Statusln": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"statusln"},
				kind:       Status,
			}},
			fn: func(v Verbose) {
				v.Statusln("statusln")
			},
		},
		"Status": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"status"},
				kind:       Status,
			}},
			fn: func(v Verbose) {
				v.Status("status")
			},
		},
		"Warningf": {
			expected: []Message{{
				format:     "hello %s",
				formatArgs: []interface{}{"warningf"},
				kind:       Warning,
			}},
			fn: func(v Verbose) {
				v.Warningf("hello %s", "warningf")
			},
		},
		"Warningln": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"warningln"},
				kind:       Warning,
			}},
			fn: func(v Verbose) {
				v.Warningln("warningln")
			},
		},
		"Warning": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"warning"},
				kind:       Warning,
			}},
			fn: func(v Verbose) {
				v.Warning("warning")
			},
		},
		"Deprecationf": {
			expected: []Message{{
				format:     "hello %s",
				formatArgs: []interface{}{"deprecationf"},
				kind:       Deprecation,
			}},
			fn: func(v Verbose) {
				v.Deprecationf("hello %s", "deprecationf")
			},
		},
		"Deprecationln": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"deprecationln"},
				kind:       Deprecation,
			}},
			fn: func(v Verbose) {
				v.Deprecationln("deprecationln")
			},
		},
		"Deprecation": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"deprecation"},
				kind:       Deprecation,
			}},
			fn: func(v Verbose) {
				v.Deprecation("deprecation")
			},
		},
		"Errorf": {
			expected: []Message{{
				format:     "hello %s",
				formatArgs: []interface{}{"errorf"},
				kind:       Error,
			}},
			fn: func(v Verbose) {
				v.Errorf("hello %s", "errorf")
			},
		},
		"Errorln": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"errorln"},
				kind:       Error,
			}},
			fn: func(v Verbose) {
				v.Errorln("errorln")
			},
		},
		"Error": {
			expected: []Message{{
				format:     "%v",
				formatArgs: []interface{}{"error"},
				kind:       Error,
			}},
			fn: func(v Verbose) {
				v.Error("error")
			},
		},
		"misc: verbosity": {
			verbosity: 2,
			expected: []Message{{
				format:    "hello",
				kind:      Status,
				verbosity: 2,
			}},
			fn: func(v Verbose) {
				v.Statusf("hello")
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := &testPrinter{}
			v := Verbose{verbosity: tc.verbosity, printer: p, enabled: true}

			tc.fn(v)

			// Check that the collected messages are expected.
			diff := cmp.Diff(tc.expected, p.messages, cmp.AllowUnexported(Message{}))
			if diff != "" {
				t.Log("did not find expected Message; want -> -, got -> +")
				t.Fatalf(diff)
			}
		})
	}
}

func pipedMessage(text string, kind MessageKind) Message {
	return Message{
		format:     "%s",
		formatArgs: []interface{}{text},
		kind:       kind,
	}
}

func TestVerboseAsWriter(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected []Message
		fn       func(v Verbose) io.Writer
	}{
		"AsWriter(Info)": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Info) },
			input: []string{"hello\n"},
			expected: []Message{
				pipedMessage("hello", Info),
			},
		},
		"AsWriter(Status)": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Status) },
			input: []string{"hello\n"},
			expected: []Message{
				pipedMessage("hello", Status),
			},
		},
		"AsWriter(Warning)": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Warning) },
			input: []string{"hello\n"},
			expected: []Message{
				pipedMessage("hello", Warning),
			},
		},
		"AsWriter(Deprecation)": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Deprecation) },
			input: []string{"hello\n"},
			expected: []Message{
				pipedMessage("hello", Deprecation),
			},
		},
		"AsWriter(Error)": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Error) },
			input: []string{"hello\n"},
			expected: []Message{
				pipedMessage("hello", Error),
			},
		},
		"Multiple Messages Single Input": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Status) },
			input: []string{"hello\nworld\n"},
			expected: []Message{
				pipedMessage("hello", Status),
				pipedMessage("world", Status),
			},
		},
		"Multiple Input": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Status) },
			input: []string{"hello\n", "world\n"},
			expected: []Message{
				pipedMessage("hello", Status),
				pipedMessage("world", Status),
			},
		},
		"Buffered Input": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Status) },
			input: []string{"hello ", "world\n"},
			expected: []Message{
				pipedMessage("hello world", Status),
			},
		},
		"Windows CRLF": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Status) },
			input: []string{"hello world\r\n"},
			expected: []Message{
				pipedMessage("hello world", Status),
			},
		},
		"Buffered Input Windows CRLF": {
			fn:    func(v Verbose) io.Writer { return v.AsWriter(Status) },
			input: []string{"hello world\r", "\n"},
			expected: []Message{
				pipedMessage("hello world", Status),
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := &testPrinter{}
			v := Verbose{printer: p, enabled: true}

			writer := tc.fn(v)
			for _, str := range tc.input {
				inputBytes := []byte(str)
				read, _ := writer.Write(inputBytes)
				if read != len(inputBytes) {
					t.Fatalf("number of bytes writter != number of bytes read")
				}
			}

			// Check that the collected messages are expected.
			diff := cmp.Diff(tc.expected, p.messages, cmp.AllowUnexported(Message{}))
			if diff != "" {
				t.Log("did not find expected Message; want -> -, got -> +")
				t.Fatalf(diff)
			}
		})
	}
}
