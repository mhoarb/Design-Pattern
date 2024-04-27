package main

import "fmt"

func main() {
	pizza := &VeggieMania{price: 20}
	//pizzaWithCheese := &CheeseTopping{pizza: pizza}

	pizzaWithTomatoAndCheese := &TomatoTopping{pizza: pizza}

	fmt.Printf("price of veggeMania with tomato and cheese topping is %d\n", pizzaWithTomatoAndCheese.getPrice())

}
