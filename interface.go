package main

import (
	"fmt"
)

type ISpeaker interface {
	Speak()
}

type SimpleSpeaker struct {
	Message string
}

func (speaker *SimpleSpeaker) Speak() {
	fmt.Println("I am speaking?", speaker.Message)
}

func main() {
	var speaker ISpeaker
	speaker = &SimpleSpeaker{"Hello"}
	speaker.Speak()
}

//output
// I am speaking? Hello
