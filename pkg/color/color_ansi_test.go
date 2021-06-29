// +build !windows

package color

import (
	"strings"
	"testing"
)

func TestAppendAnsiParameter(t *testing.T) {
	var builder strings.Builder

	appendAnsiParameter(&builder, "0")
	if builder.String() != "\x1B[0" {
		t.Fatalf("expected: \"^[0\", got: %#v", builder.String())
	}

	appendAnsiParameter(&builder, "31")
	if builder.String() != "\x1B[0;31" {
		t.Fatalf("expected: \"^[0;31\", got: %#v", builder.String())
	}
}

func TestApply(t *testing.T) {
	tests := map[string]struct {
		expected string
		input    Style
	}{
		"Plain": {
			input:    Plain(),
			expected: "test",
		},
		"Foreground": {
			input:    Plain().Foreground(Red),
			expected: "\x1B[31m" + "test" + ansiReset,
		},
		"Background": {
			input:    Plain().Background(Red),
			expected: "\x1B[41m" + "test" + ansiReset,
		},
		"Bold": {
			input:    Plain().Bold(true),
			expected: "\x1B[1m" + "test" + ansiReset,
		},
		"All": {
			input:    Plain().Foreground(Green).Background(Red).Bold(true),
			expected: "\x1B[1;32;41m" + "test" + ansiReset,
		},
		"Color Red": {
			input:    Plain().Foreground(Red),
			expected: "\x1B[31m" + "test" + ansiReset,
		},
		"Color Green": {
			input:    Plain().Foreground(Green),
			expected: "\x1B[32m" + "test" + ansiReset,
		},
		"Color Yellow": {
			input:    Plain().Foreground(Yellow),
			expected: "\x1B[33m" + "test" + ansiReset,
		},
		"Color Blue": {
			input:    Plain().Foreground(Blue),
			expected: "\x1B[34m" + "test" + ansiReset,
		},
		"Color Magenta": {
			input:    Plain().Foreground(Magenta),
			expected: "\x1B[35m" + "test" + ansiReset,
		},
		"Color Cyan": {
			input:    Plain().Foreground(Cyan),
			expected: "\x1B[36m" + "test" + ansiReset,
		},
		"Color White": {
			input:    Plain().Foreground(White),
			expected: "\x1B[39m" + "test" + ansiReset,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.Apply("test")
			if tc.expected != got {
				t.Fatalf("expected: %#v, got: %#v", tc.expected, got)
			}
		})
	}
}
