package main 
type PeperoniPizza struct {
	name string 
	sauce string 
	toppings [] string
}

func(p *PeperoniPizza)getName() string {
	return p.name
}
func(p *PeperoniPizza)getSauce() string {
	return p.sauce
}
func (p *PeperoniPizza) getToppings()[]string {
	return p.toppings
}
func(p * PeperoniPizza) setToppings(topping []string) {
	p.toppings = topping
}
