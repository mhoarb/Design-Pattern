package main

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		observerList: nil,
		name:         "product",
		inStock:      false,
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}
func(i * Item)deRegister(o Observer) {
	i.observerList =
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func RemoveFromSlice(observerList []Observer , observerToRemove Observer)[]Observer {
	observerListLength := len(observerList)
	for i , observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[]
		}
	}
}
