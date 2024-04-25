package main

import "fmt"

var _ Printer = &HP{}

type HP struct {
}

func (H *HP) printFile() {
	fmt.Println("print by Hp printer")
}
