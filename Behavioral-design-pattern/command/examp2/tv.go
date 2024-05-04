package main

import "fmt"

type TV struct {
	name      string
	isRunning bool
}

func (t *TV) on() {
	if t.isRunning == true {
		fmt.Printf("this device: %s is already on", t.name)
		return
	}
	t.isRunning = true
	fmt.Println("turning tv on")
}

func (t *TV) off() {
	if t.isRunning == false {
		fmt.Printf("this device : %s is already off", t.name)
		return
	}
	t.isRunning = false
	fmt.Println("turning tv off")
}
