package main

import "fmt"


type HpAdapter struct {
	printer *HpPrinter
}

func(ha *HpAdapter)PrintA4() {
	fmt.Println("adapting to A4 size for HP printer")
}