package main 

func(d *DefaultPizzaBuilder)BuildMargheritaPizza() IPizza {
	newPizza := d.margheritoPizza.Clone()

	return newPizza
}
//........