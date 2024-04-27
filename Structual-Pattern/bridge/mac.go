package main

import "fmt"

var _ Computer = &Mac{}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}
