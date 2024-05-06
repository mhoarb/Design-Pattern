package main

import "fmt"

func main() {
	caretaker := &Caretaker{mementoArray: make([]*Memento, 0)}

	originator := &Originator{state: "A"}

	fmt.Printf("orginator current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("orginator current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
