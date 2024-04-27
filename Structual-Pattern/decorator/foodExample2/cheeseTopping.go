package main

import "fmt"

var _ Pizza = &CheeseTopping{}

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() int {
	fmt.Println("cheese topping ...")
	return c.pizza.getPrice() + 150
}
