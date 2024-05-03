package main

import "fmt"

var _ Departmant = &Doctor{}

type Doctor struct {
	next Departmant
}

func (d *Doctor) execute(patient *Patient) {
	if patient.doctorCheckupDone {
		fmt.Println("Doctor checkUp already done")
		d.next.execute(patient)
	}
	fmt.Println("Doctor checking patient")
	patient.doctorCheckupDone = true
	d.next.execute(patient)
}

func (d *Doctor) setNext(next Departmant) {
	d.next = next
}
