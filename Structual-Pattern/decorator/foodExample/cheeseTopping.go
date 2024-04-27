package main

var _ Pizza = &CheeseTopping{}

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 100
}
