package main

func main() {
	StationManager := newStatrionManager()
	PassengerTrain := &PassengerTrain{
		mediator: StationManager,
	}
	freeightTrain := &FreightTrain{mediator: StationManager}

	PassengerTrain.arrive()
	freeightTrain.arrive()
	PassengerTrain.depart()
}
