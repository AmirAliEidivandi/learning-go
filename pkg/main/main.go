package main

import (
	"learning-go/display"
	"learning-go/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("an exciting message")
}