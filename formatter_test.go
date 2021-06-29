package clout

import (
	"fmt"
	"testing"
)

// testHighlighter is an implementation of Highlighter that "highlights" values with brackets.
type testHighlighter struct {
	value interface{}
}

func (t testHighlighter) Value() interface{} {
	return t.value
}

func (t testHighlighter) Apply(str string) string {
	return fmt.Sprintf("{%s}", str)
}

func TestFormatText(t *testing.T) {
	tests := map[string]struct {
		message  Message
		colors   bool
		expected string
	}{
		"Apply Colors": {
			expected: "hello {world}",
			colors:   true,
			message: Message{
				format:     "hello %s",
				formatArgs: []interface{}{testHighlighter{value: "world"}},
			},
		},
		"Discard Colors": {
			expected: "hello world",
			colors:   false,
			message: Message{
				format:     "hello %s",
				formatArgs: []interface{}{testHighlighter{value: "world"}},
			},
		},
		"Apply Colors Without Highlight": {
			expected: "hello world",
			colors:   true,
			message: Message{
				format:     "hello %s",
				formatArgs: []interface{}{"world"},
			},
		},
		"Apply Colors With Debug Verb": {
			expected: "hello {\"world\"}",
			colors:   true,
			message: Message{
				format:     "hello %#v",
				formatArgs: []interface{}{testHighlighter{value: "world"}},
			},
		},
		"Discard Colors With Debug Verb": {
			expected: "hello \"world\"",
			colors:   false,
			message: Message{
				format:     "hello %#v",
				formatArgs: []interface{}{testHighlighter{value: "world"}},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := formatText(&tc.message, tc.colors)

			if tc.expected != got {
				t.Fatalf("expected: %s, got: %s", tc.expected, got)
			}
		})
	}
}
