package main

import "fmt"

var _ Train = &PassengerTrain{}

type PassengerTrain struct {
	mediator Mediator
}

func (p *PassengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("passengerTrain: arrival blocked ,waiting.., ")
		return
	}
	fmt.Println("passengerTrain : Arrived")
}

func (p *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: leaving...")
	p.mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain : Arrival permitted , arriving ...")
	p.arrive()
}
