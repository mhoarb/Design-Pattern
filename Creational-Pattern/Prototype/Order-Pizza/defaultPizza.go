package main 

type DefaultPizzaBuilder struct{
	margheritoPizza *MargheritoPizza
	peperoniPizza *PeperoniPizza
}

func NewDefaultPizza() *DefaultPizzaBuilder{
	return &DefaultPizzaBuilder{
		margheritoPizza: &MargheritoPizza{
			name: "Margherita",
			sauce:"tomato",
			toppings: []string {"mozzarella" , "basil"},
		},
		peperoniPizza: &PeperoniPizza{
			name: "Peperoni",
			sauce: "Tomato",
			toppings: []string {"mozzarella" , "Peperoni"},
		},
	}
}