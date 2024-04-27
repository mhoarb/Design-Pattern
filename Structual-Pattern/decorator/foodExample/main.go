package main

import "fmt"

func main() {
	pizza := &VeggieMania{price: 100}

	pizzaWithCheese := &CheeseTopping{pizza: pizza}

	pizzaWithTomatoTopping := &TomatoTopping{pizza: pizza}

	fmt.Printf("price of veggmania with tomato %d\n price with cheese topping %d",
		pizzaWithCheese.getPrice(), pizzaWithTomatoTopping.getPrice())
}
