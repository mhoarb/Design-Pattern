package main

import "fmt"

var _ Computer = &WindowsAdapter{}

type WindowsAdapter struct {
	windwsMachin *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windwsMachin.InsertIntoLightningPort()
}
