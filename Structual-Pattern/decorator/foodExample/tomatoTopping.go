package main

var _ Pizza = &TomatoTopping{}

type TomatoTopping struct {
	pizza Pizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 10
}
