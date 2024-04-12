package main 

func(p *MargheritoPizza)Clone() IPizza {
	newPizza := &MargheritoPizza{
		name: p.name,
		sauce: p.sauce,
		toppings: make([]string,len(p.toppings) ),
	}
	copy(newPizza.toppings , p.toppings)
	return newPizza
}

func(p *PeperoniPizza) Clone() IPizza {
	newPizza := &PeperoniPizza{
		name: p.name,
		sauce: p.sauce,
		toppings: make([]string,len(p.toppings) ),
	}
	copy(newPizza.toppings, p.toppings)
	return newPizza
}
