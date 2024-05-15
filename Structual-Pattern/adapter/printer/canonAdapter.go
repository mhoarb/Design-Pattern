package main

import "fmt"

type CanonAdapter struct {
	printer *CanonPrinter
}


func(ca *CanonAdapter)PrintA4 () {

	fmt.Println("adapting to A4 size for canon printer")

	ca.printer.PrintWithHPCanon()

}
