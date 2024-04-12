package main 
type IPizza interface {
	setToppings([]string)

	getToppings() []string
	getName() string
	getSauce()string
}

