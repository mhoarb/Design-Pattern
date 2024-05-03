package main

import "fmt"

var _ Departmant = &Doctor{}

type Doctor struct {
	next Departmant
}

func (d *Doctor) execute(patient *Patient) {
	if patient.doctorCheckUpDone {
		fmt.Println("doctor check already done")
		d.next.execute(patient)
		return
	}
	fmt.Println("doctor Checking patient")
	patient.doctorCheckUpDone = true
	d.next.execute(patient)
}

func (d *Doctor) setNext(next Departmant) {
	d.next = next
}
