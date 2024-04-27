package main

import "fmt"

var _ Printer = &Epson{}

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("printing bu a EPSON printer")
}
