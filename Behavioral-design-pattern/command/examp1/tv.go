package main

import "fmt"

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	fmt.Println("turning TV on")
}

func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("turning TV off")
}
