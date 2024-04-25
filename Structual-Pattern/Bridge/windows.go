package main

import "fmt"

var _ Computer = &Windows{}

type Windows struct {
	printer Printer
}

func (w *Windows) print() {
	fmt.Println("print the request for windows")
	w.printer.printFile()
}

func (w *Windows) setPrinter(printer Printer) {
	w.printer = printer
}
