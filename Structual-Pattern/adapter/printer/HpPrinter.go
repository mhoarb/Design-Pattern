package main

import "fmt"

type HpPrinter struct {

}

func(h *HpPrinter) PrintWithHP() {
	fmt.Println("print with Hp printer")
}
