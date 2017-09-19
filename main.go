package main

import (
	"dependency"
	"speaker"
)

func main() {
	spk1 := speaker.NewSpeaker()
	spk1.Speak()
	// Outputs: Hello universe

	spk2 := dependency.NewSpeaker()
	spk2.Speak()
	// Outputs: Hello alternate universe

	// spk1 = spk2
	// ./main.go:17:8: cannot use spk2 (type *"splithorizon/vendor/dependency/vendor/speaker".Speaker) as type *"splithorizon/vendor/speaker".Speaker in assignment
}
