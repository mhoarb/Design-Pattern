package main

import "fmt"

var _ Departmant = &Medical{}

type Medical struct {
	next Departmant
}

func (m *Medical) execute(patient *Patient) {
	if patient.medicineDone {
		fmt.Println("medicine already given to patient")
		m.next.execute(patient)
	}
	fmt.Println("medical giving medicine to patient")
	patient.medicineDone = true
	m.next.execute(patient)
}

func (m *Medical) setNext(next Departmant) {
	m.next = next
}
