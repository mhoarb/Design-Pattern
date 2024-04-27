package main

import "fmt"

var _ Computer = &Windows{}

type Windows struct {
	name string
}

func (w *Windows) InsertIntoLightningPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
