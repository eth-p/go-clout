package fitm

import (
	"fmt"
)

// Verb is a Go fmt formatting verb.
// See https://golang.org/pkg/fmt/#hdr-Printing for details.
type Verb struct {
	flags string
	verb  string
}

// String returns the fmt.Printf representation of the Verb.
func (v Verb) String() string {
	return "%" + v.flags + v.verb
}

// Format formats a value with the Verb, similarly to fmt.Sprintf.
func (v Verb) Format(with interface{}) string {
	return fmt.Sprintf(v.String(), with)
}

// NewVerb creates a new Verb.
func NewVerb(verb string) Verb {
	return Verb{
		flags: "",
		verb:  verb,
	}
}

// NewVerbWithFlags creates a new Verb with flags.
func NewVerbWithFlags(verb string, flags string) Verb {
	return Verb{
		flags: flags,
		verb:  verb,
	}
}

// errorString returns the fmt.Printf "error" representation of the Verb.
// Examples:
//
//   %!s(MISSING)
//
func (v Verb) errorString(err string) string {
	return "%" + "!" + v.flags + v.verb + "(" + err + ")"
}

// VerbOrLiteral is either a Verb, or a literal string.
type VerbOrLiteral interface{}
