package main

import "fmt"

var _ Computer = &Mac{}

type Mac struct {
	printer Printer
}

func (m *Mac) print() {
	fmt.Println("print the request for mac")
	m.printer.printFile()
}

func (m *Mac) setPrinter(printer Printer) {
	m.printer = printer
}
