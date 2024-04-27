package main

import "fmt"

var _ Pizza = &TomatoTopping{}

type TomatoTopping struct {
	pizza Pizza
}

func (t *TomatoTopping) getPrice() int {
	fmt.Println("tomato topping ...")
	return t.pizza.getPrice() + 100
}
