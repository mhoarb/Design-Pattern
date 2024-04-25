package main

import "fmt"

var _ Computer = &Mac{}

type Mac struct{}

func (m *Mac) insertInToLightningPort() {
	fmt.Println(" lightning connector is plugged into mac machine ")

}
