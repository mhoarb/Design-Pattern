package main

type Printer struct {
}

func (p *Printer) print(printer IPrint) {
	printer.PrintA4()
}
