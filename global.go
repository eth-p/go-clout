package clout

import (
	"sync"
)

var globalPrinterMutex sync.RWMutex
var globalPrinter PrinterInterface
var globalVerbosityMutex sync.RWMutex
var globalVerbosity MessageVerbosity

// GetPrinter gets the global PrinterInterface instance.
func GetPrinter() PrinterInterface {
	globalPrinterMutex.RLock()
	defer globalPrinterMutex.RUnlock()
	return globalPrinter
}

// SetPrinter sets the global PrinterInterface instance.
func SetPrinter(processor PrinterInterface) {
	globalPrinterMutex.Lock()
	globalPrinter = processor
	globalPrinterMutex.Unlock()
}

// GetVerbosity gets the minimum MessageVerbosity required for messages to be displayed.
func GetVerbosity() MessageVerbosity {
	globalVerbosityMutex.RLock()
	defer globalVerbosityMutex.RUnlock()
	return globalVerbosity
}

// SetVerbosity sets the minimum MessageVerbosity required for messages to be displayed.
// The higher the number, the more verbose the output will become.
func SetVerbosity(verbosity MessageVerbosity) {
	globalVerbosityMutex.Lock()
	globalVerbosity = verbosity
	globalVerbosityMutex.Unlock()
}

// V creates a struct to print messages.
//
// Example:
//
//     V(2).Warningf("unknown path: %v", highlight.Cyan("/not-a-path"))
func V(verbosity MessageVerbosity) *Verbose {
	return &Verbose{
		enabled:   verbosity <= GetVerbosity(),
		verbosity: verbosity,
		printer:   GetPrinter(),
	}
}

func init() {
	SetPrinter(NewPrinterWithDefaults(true))
	SetVerbosity(defaultVerbosity)
}
