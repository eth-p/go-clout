package color

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConstructors(t *testing.T) {
	tests := map[string]struct {
		expected Style
		got      Style
	}{
		"Plain": {
			got: Plain(),
			expected: Style{
				foreground: None,
				background: None,
				bold:       false,
			},
		},
		"Foreground": {
			got: Foreground(999),
			expected: Style{
				foreground: 999,
				background: None,
				bold:       false,
			},
		},
		"Background": {
			got: Background(999),
			expected: Style{
				foreground: None,
				background: 999,
				bold:       false,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			diff := cmp.Diff(tc.expected, tc.got, cmp.AllowUnexported(Style{}))
			if diff != "" {
				t.Log("did not find expected Style; want -> -, got -> +")
				t.Fatalf(diff)
			}
		})
	}
}

func TestFunctions(t *testing.T) {
	tests := map[string]struct {
		expected Style
		got      Style
	}{
		"Foreground": {
			got: Plain().Foreground(999),
			expected: Style{
				foreground: 999,
				background: None,
				bold:       false,
			},
		},
		"Background": {
			got: Plain().Background(999),
			expected: Style{
				foreground: None,
				background: 999,
				bold:       false,
			},
		},
		"Bold": {
			got: Plain().Bold(true),
			expected: Style{
				foreground: None,
				background: None,
				bold:       true,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			diff := cmp.Diff(tc.expected, tc.got, cmp.AllowUnexported(Style{}))
			if diff != "" {
				t.Log("did not find expected Style; want -> -, got -> +")
				t.Fatalf(diff)
			}
		})
	}
}
