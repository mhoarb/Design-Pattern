package main

import "fmt"

var _ Computer = &WindowsAdapter{}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) insertInToLightningPort() {
	fmt.Println("adapter converts lightning signal to usb")
	w.windowsMachine.insertInToLightningPort()
}
