package speaker

import "fmt"

type Speaker struct{}

func (s *Speaker) Speak() {
	fmt.Println("Hello universe")
}

func NewSpeaker() *Speaker {
	return &Speaker{}
}