package main

import "fmt"

type Medical struct {
	next Departmant
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.paymentDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Departmant) {
	m.next = next
}
