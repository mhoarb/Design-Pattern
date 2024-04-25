package main

import "fmt"

var _ Computer = &Windows{}

type Windows struct{}

func (w *Windows) insertInToLightningPort() {
	fmt.Println("usb connector is plugged into windows machine")
}
