package speaker

import "fmt"

type Speaker struct{}

func (s Speaker) Speak() {
	fmt.Println("Hello alternate universe")
}

func NewSpeaker() *Speaker {
	return &Speaker{}
}
