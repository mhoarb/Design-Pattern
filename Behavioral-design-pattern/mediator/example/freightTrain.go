package main

import "fmt"

var _ Train = &FreightTrain{}

type FreightTrain struct {
	mediator Mediator
}

func (f FreightTrain) arrive() {
	if !f.mediator.canArrive(f) {
		fmt.Println("PassengerTrain : Arrival blocked, waiting ...")
		return
	}
	fmt.Println("PassengerTrain : Arrived")
}

func (f FreightTrain) depart() {
	fmt.Println("PassengerTrain : leaving")
	f.mediator.notifyAboutDeparture()
}

func (f FreightTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted")
	f.arrive()
}
