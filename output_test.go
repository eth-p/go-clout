package clout

import (
	"bytes"
	"testing"

	"github.com/eth-p/clout/pkg/color"
)

func TestOutput(t *testing.T) {
	tests := map[string]struct {
		message  Message
		expected string
		init     func(output Output) Output
	}{
		"With Prefix": {
			expected: "\x1B[1;31merror:\u001B[0m \x1B[31mhello world\x1B[0m\n",
			message:  New(Info, 2, "hello world"),
			init: func(output Output) Output {
				return output.
					WithColor(color.Foreground(color.Red)).
					WithPrefix("error:", color.Foreground(color.Red).Bold(true)).
					WithColors(true)
			},
		},
		"With Colors": {
			expected: "\x1B[31mhello world\x1B[0m\n",
			message:  New(Info, 2, "hello world"),
			init: func(output Output) Output {
				return output.
					WithColor(color.Foreground(color.Red)).
					WithColors(true)
			},
		},
		"Without Colors": {
			expected: "error: hello world\n",
			message:  New(Info, 2, "hello world"),
			init: func(output Output) Output {
				return output.
					WithColor(color.Foreground(color.Red)).
					WithPrefix("error:", color.Plain()).
					WithColors(false)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			buf := new(bytes.Buffer)

			output := tc.init(OutputFromWriter(buf))
			if err := output.write(&tc.message); err != nil {
				t.Fatal(err)
			}

			got := buf.String()
			if tc.expected != got {
				t.Fatalf("expected: %#v, got: %#v", tc.expected, got)
			}
		})
	}
}
