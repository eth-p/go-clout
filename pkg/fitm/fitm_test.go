package fitm

import (
	"testing"
)

func TestSprintf(t *testing.T) {
	tests := map[string]struct {
		format   string
		args     []interface{}
		expected string
		fn       FormatMitm
	}{
		"no change": {
			format:   "hello, %s",
			args:     []interface{}{"world"},
			expected: "hello, world",

			fn: func(verb Verb, val interface{}) (Verb, interface{}) {
				return verb, val // no change
			},
		},
		"replace args": {
			format:   "hello, %s! I'm %s",
			args:     []interface{}{"world", "ethan"},
			expected: "hello, mitm! I'm mitm",

			fn: func(verb Verb, val interface{}) (Verb, interface{}) {
				return verb, "mitm"
			},
		},
		"replace verbs": {
			format:   "%v %v",
			args:     []interface{}{1.001, 2.0},
			expected: "1.00 2.00",

			fn: func(verb Verb, val interface{}) (Verb, interface{}) {
				return Verb{flags: "1.2", verb: "f"}, val
			},
		},
		"wrap args": {
			format:   "the number %1.3f is %1.0f when truncated",
			args:     []interface{}{1.123, 1.123},
			expected: "the number (1.123) is (1) when truncated",

			fn: func(verb Verb, val interface{}) (Verb, interface{}) {
				return Preformatted("(" + verb.Format(val) + ")")
			},
		},
		"too few args": {
			format:   "%s %s",
			args:     []interface{}{"one"},
			expected: "one %!s(MISSING)",

			fn: func(verb Verb, val interface{}) (Verb, interface{}) {
				return verb, val // no change
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Sprintf(tc.fn, tc.format, tc.args...)
			if got != tc.expected {
				t.Fatalf("expected '%v', got '%v'", tc.expected, got)
			}
		})
	}
}
