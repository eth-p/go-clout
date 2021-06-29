// +build windows

package color

// Apply applies text styling to a string.
// On Windows, this isn't supported for compatibility reasons.
//
// ANSI Color codes in conhost.exe (the default terminal emulator) requires Windows 10 and an explicit call to
// Kernel32.dll!SetConsoleMode with the ENABLE_VIRTUAL_TERMINAL_INPUT flag.
//
// https://docs.microsoft.com/en-us/windows/console/setconsolemode
func (s Style) Apply(str string) string {
	return str
}
