package fitm

import (
	"fmt"
	"strings"
)

// Parse parses a Go fmt string into words.
func Parse(format string) ([]VerbOrLiteral, error) {
	var words []VerbOrLiteral

	offset := 0
	index := strings.Index(format[offset:], "%")

	for index != -1 {
		// Append the literal text before the verb.
		beforeText := format[offset:(offset + index)]
		if beforeText != "" {
			words = append(words, beforeText)
		}
		offset += index

		// Parse the verb.
		verb, err := parseVerb(format[offset:])
		if err != nil {
			return nil, err
		}

		// Append the verb and update the offset to skip past the parsed verb.
		words = append(words, verb)
		offset += len("%") + len(verb.flags) + len(verb.verb)

		// Look for the next occurrence.
		index = strings.Index(format[offset:], "%")
	}

	// Append any remaining literal text.
	remaining := format[offset:]
	if remaining != "" {
		words = append(words, remaining)
	}

	// Return the parsed data.
	return words, nil
}

// parseVerb attempts to parse a formatting verb.
// This isn't by any means accurate, but it should work for the purposes of mitm'ing a format string.
//
// Pattern:
//   %([^A-Za-z%]*)([A-Za-z%])
//     ~~~~~~~~~~~  ~~~~~~~~~
//     ^            ^
//   flags        verb
func parseVerb(verb string) (Verb, error) {
	inputVerb := verb
	if verb[0] != '%' {
		return Verb{}, fmt.Errorf("format verb does not start with '%%': %v", inputVerb)
	}

	// Remove the leading %.
	verb = verb[len("%"):]

	// Find index of the verb character after the flags.
	index := strings.IndexFunc(verb, func(r rune) bool {
		return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || r == '%'
	})

	if index == -1 {
		return Verb{}, fmt.Errorf("format verb not valid: %v", inputVerb)
	}

	// Parse out the flags and the verb character.
	verbFlags := verb[:index]
	verbChar := verb[index:(index + 1)]

	// Return the verb.
	return Verb{
		flags: verbFlags,
		verb:  verbChar,
	}, nil
}
