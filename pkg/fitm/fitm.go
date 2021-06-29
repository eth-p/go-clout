package fitm

import (
	"fmt"
	"io"
	"strings"
)

// FormatMitm is a function that gets called for each verb in fmt format string.
//
// This function is called with the original value of a verb+value tuple, and can be used to replace either (or both)
// the formatting verb and formatted value.
//
// Example:
//
//   func allDebug(verb Verb, val interface{}) (Verb, interface{}) {
//      return NewVerbWithFlags("v", "#"), val
//   }
//
//   fitm.Printf(allDebug, "hello %s", "world") // -> fmt.Printf("hello %#v", "world")
type FormatMitm func(verb Verb, val interface{}) (Verb, interface{})

// Preformatted can be used inside a FormatMitm function to return pre-formatted text.
// This is useful for when you replace a formatting verb with a literal value.
func Preformatted(str string) (Verb, interface{}) {
	return NewVerb("s"), str
}

// Sprintf is a version of fmt.Sprintf that accepts a FormatMitm function.
func Sprintf(mitm FormatMitm, format string, a ...interface{}) string {
	// Parse the verbs.
	newFormat, newArgs, err := fmtMitm(mitm, format, a)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(newFormat, newArgs...)
}

// Printf is a version of fmt.Printf that accepts a FormatMitm function.
func Printf(mitm FormatMitm, format string, a ...interface{}) {
	// Parse the verbs.
	newFormat, newArgs, err := fmtMitm(mitm, format, a)
	if err != nil {
		panic(err)
	}

	fmt.Printf(newFormat, newArgs...)
}

// Fprintf is a version of fmt.Fprintf that accepts a FormatMitm function.
func Fprintf(mitm FormatMitm, writer io.Writer, format string, a ...interface{}) (int, error) {
	// Parse the verbs.
	newFormat, newArgs, err := fmtMitm(mitm, format, a)
	if err != nil {
		panic(err)
	}

	return fmt.Fprintf(writer, newFormat, newArgs...)
}

// fmtMitm applies a FormatMitm function to a format string and its corresponding arguments.
// This allows for formatting verbs (or arguments) to be conditionally replaced at runtime.
func fmtMitm(mitm FormatMitm, format string, args []interface{}) (string, []interface{}, error) {
	// Parse the format string into an array of verbs and string literals.
	parsed, err := Parse(format)
	if err != nil {
		return "", nil, err
	}

	// Create a new formatting string from all the parsed items.
	var argIndex = 0
	var argLast = len(args)
	var newFormat strings.Builder
	newArguments := make([]interface{}, len(args))

	for _, item := range parsed {
		if str, ok := item.(string); ok {
			// If it's a string literal, use it directly.
			newFormat.WriteString(str)
		} else if verb, ok := item.(Verb); ok {
			// If it's a verb, get the corresponding argument.
			if argIndex >= argLast {
				newFormat.WriteString("%")
				newFormat.WriteString(verb.errorString("MISSING"))
				continue
			}

			arg := args[argIndex]

			// Apply the mitm function to change the verb and/or its corresponding argument.
			newVerb, newValue := mitm(verb, arg)

			newFormat.WriteString(newVerb.String())
			newArguments[argIndex] = newValue

			argIndex += 1
		} else {
			panic("VerbOrLiteral is neither Verb nor string")
		}
	}

	// Reconstruct a formatting string and run it through fmt.Sprintf.
	return newFormat.String(), newArguments, nil
}
