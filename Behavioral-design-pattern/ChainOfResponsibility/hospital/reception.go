package main

import "fmt"

var _ Departmant = &Reception{}

type Reception struct {
	next Departmant
}

func (r *Reception) execute(patient *Patient) {
	if patient.registerDone {
		fmt.Println("patent registration is already done")
		r.next.execute(patient)
	}
	fmt.Println("reception registering patient")
	patient.registerDone = true
	r.next.execute(patient)

}

func (r *Reception) setNext(next Departmant) {
	r.next = next
}
