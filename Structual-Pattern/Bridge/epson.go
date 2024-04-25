package main

import "fmt"

var _ Printer = &Epson{}

type Epson struct{}

func (e *Epson) printFile() {
	fmt.Println("printing by Epson printer")
}
