package main

import "fmt"

var _ Departmant = &Reception{}

type Reception struct {
	next Departmant
}

func (r *Reception) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Patient regestration already done")
		r.next.execute(patient)
		return
	}
	fmt.Println("Reception regestring patient")
	patient.registrationDone = true
	r.next.execute(patient)
}

func (r *Reception) setNext(next Departmant) {
	r.next = next
}
