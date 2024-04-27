package main

import "fmt"

var _ Computer = &Windows{}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}
