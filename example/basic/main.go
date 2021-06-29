package main

import (
	"github.com/eth-p/clout"
)

var animals = []string{
	"shark",
	"cat",
	"dog",
}

var pets = map[string]bool{
	"cat": true,
	"dog": true,
}

func isPet(animal string) bool {
	isPet, inPets := pets[animal]
	return inPets && isPet
}

func main() {

	// Set the verbosity level to 3, to show all the messages in this example.
	// The default level is 2.
	clout.SetVerbosity(3)

	// Info messages are for when you want to inform the user about what your program is doing.
	clout.V(2).Infof("Checking if your animals are pets...")

	notPets := 0
	for _, animal := range animals {
		// Status messages are for when you want to update the user on how the program is progressing.
		clout.V(3).Statusf("Checking animal %#v...", animal)
		if !isPet(animal) {
			clout.V(2).Warningf("A %s is not a pet.", animal)
			notPets++
		}
	}

	if notPets > 0 {
		clout.V(1).Errorf("%d of your pets are actually wild animals.", notPets)
	}

}
