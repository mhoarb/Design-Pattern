package main

import "fmt"

var _ Printer = &HP{}

type HP struct {
}

func (H *HP) PrintFile() {
	fmt.Println("printing by a epson Printer")
}
