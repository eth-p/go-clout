package fitm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseVerb(t *testing.T) {
	tests := []struct {
		input         string
		expectedVerb  string
		expectedFlags string
		shouldError   bool
	}{
		{
			input:         "%%",
			expectedVerb:  "%",
			expectedFlags: "",
		},
		{
			input:         "%v",
			expectedVerb:  "v",
			expectedFlags: "",
		},
		{
			input:         "%#v",
			expectedVerb:  "v",
			expectedFlags: "#",
		},
		{
			input:         "%9.2f",
			expectedVerb:  "f",
			expectedFlags: "9.2",
		},
		{
			input:         "%+q",
			expectedVerb:  "q",
			expectedFlags: "+",
		},
		{
			input:         "%-n",
			expectedVerb:  "n",
			expectedFlags: "-",
		},
		{
			input:         "% s",
			expectedVerb:  "s",
			expectedFlags: " ",
		},
		{ // should not parse " leftover text"
			input:         "%+q leftover text",
			expectedVerb:  "q",
			expectedFlags: "+",
		},
		{
			input:       "%",
			shouldError: true,
		},
		{
			input:       "%00000",
			shouldError: true,
		},
		{
			input:       "abcdef",
			shouldError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			gotVerb, err := parseVerb(tc.input)

			// Check that the returned error is expected.
			if tc.shouldError && err == nil {
				t.Fatalf("expected error, got no error")
			} else if !tc.shouldError && err != nil {
				t.Fatalf("expected no error, got err: %v", err)
			} else if tc.shouldError && err != nil {
				return // All good.
			}

			// Compare verb and verb flags.
			if tc.expectedVerb != gotVerb.verb {
				t.Fatalf("expected verb: %v, got: %v", tc.expectedVerb, gotVerb.verb)
			}

			if tc.expectedFlags != gotVerb.flags {
				t.Fatalf("expected flags: %v, got: %v", tc.expectedFlags, gotVerb.flags)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input       string
		expected    []VerbOrLiteral
		shouldError bool
	}{
		"no verbs": {
			input: "hello world",
			expected: []VerbOrLiteral{
				"hello world",
			},
		},
		"verb with text before it": {
			input: "before %v",
			expected: []VerbOrLiteral{
				"before ",
				Verb{verb: "v"},
			},
		},
		"verb with text after it": {
			input: "%v after",
			expected: []VerbOrLiteral{
				Verb{verb: "v"},
				" after",
			},
		},
		"multiple_verbs": {
			input: "%v%s",
			expected: []VerbOrLiteral{
				Verb{verb: "v"},
				Verb{verb: "s"},
			},
		},
		"multiple_verbs_with_text": {
			input: "%v %s",
			expected: []VerbOrLiteral{
				Verb{verb: "v"},
				" ",
				Verb{verb: "s"},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Parse(tc.input)

			// Check that the returned error is expected.
			if tc.shouldError && err == nil {
				t.Fatalf("expected error, got no error")
			} else if !tc.shouldError && err != nil {
				t.Fatalf("expected no error, got err: %v", err)
			}

			// Compare parsed data.
			diff := cmp.Diff(tc.expected, got, cmp.AllowUnexported(Verb{}))
			if diff != "" {
				t.Log("did not find expected VerbOrLiteral; want -> -, got -> +")
				t.Fatalf(diff)
			}
		})
	}
}
