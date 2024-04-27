package main

import "fmt"

func main() {
	pizza := &VeggieMania{price: 100}

	pizzaWithCheesTopping := &CheeseTopping{pizza: pizza}
	pizzaWithTomatoTopping := &TomatoTopping{pizza: pizza}

	fmt.Printf("price pizza with cheese topping is :%d\n", pizzaWithCheesTopping.getPrice())
	fmt.Printf("price pizza with tomato topping is :%d\n", pizzaWithTomatoTopping.getPrice())

}
