package clout

import (
	"os"
	"strings"
)

// argsToFormat generates a format string for an array of arguments.
func argsToFormat(args []interface{}) string {
	argsLen := len(args)
	argsFmt := ""
	if argsLen > 0 {
		argsFmt = "%v" + strings.Repeat(" %v", argsLen-1)
	}

	return argsFmt
}

// supportsColor checks if an os.File (e.g. stdout) supports colors.
//
// This is based on the following rules:
// - If $NO_COLOR is defined, disable colors.
// - If the os.File is not a terminal, disable colors.
// - Otherwise, enable colors.
func supportsColor(fd *os.File) bool {
	// If NO_COLOR is defined, we should not be printing color.
	if _, exists := os.LookupEnv("NO_COLOR"); exists {
		return false
	}

	// If the output FD is not a terminal, we should not be printing color.
	if stat, _ := fd.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		return false
	}

	// It's fine to print color.
	return true
}
