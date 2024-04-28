package main

import "fmt"

func main() {
	game := newGame()

	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorysingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("dressType : %s\nDressColor %s\n", dressType, dress.getColor())

	}
}
