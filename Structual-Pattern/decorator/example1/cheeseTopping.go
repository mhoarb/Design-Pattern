package main

var _ Pizza = &CheeseTopping{}

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() int {
	price := c.pizza.getPrice()
	return price + 15
}
