package clout

import (
	"testing"
)

func TestArgsToFormat(t *testing.T) {
	tests := map[string]struct {
		args     []interface{}
		expected string
	}{
		"Zero": {
			expected: "",
			args:     []interface{}{},
		},
		"One": {
			expected: "%v",
			args:     []interface{}{"one"},
		},
		"Two": {
			expected: "%v %v",
			args:     []interface{}{"one", "two"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := argsToFormat(tc.args)
			if tc.expected != got {
				t.Fatalf("did not find expected formatting; want '%s', got '%s'", tc.expected, got)
			}
		})
	}
}
