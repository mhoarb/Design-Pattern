package main

var _ Pizza = &TomatoTopping{}

type TomatoTopping struct {
	pizza Pizza
}

func (t *TomatoTopping) getPrice() int {
	price := t.pizza.getPrice()
	return price + 10
}
