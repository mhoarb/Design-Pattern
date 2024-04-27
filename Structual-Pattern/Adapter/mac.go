package main

import "fmt"

var _ Computer = &Mac{}

type Mac struct {
	name string
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
