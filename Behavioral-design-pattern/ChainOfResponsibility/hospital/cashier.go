package main

import "fmt"

var _ Departmant = &Cashier{}

type Cashier struct {
	next Departmant
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("peyment already done")
		c.next.execute(p)

	}
	fmt.Println("cashier getting money from patient ")
	p.medicineDone = true
	fmt.Println("all in done")

}

func (c *Cashier) setNext(next Departmant) {
	c.next = next
}
