package main

import "fmt"

type CanonPrinter struct {
}

func (c *CanonPrinter) PrintWithHPCanon() {
	fmt.Println("print with Canon Printer")
}
