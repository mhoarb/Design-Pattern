package main

type MargheritoPizza struct {
	name     string
	sauce    string
	toppings []string
}

func (c *MargheritoPizza) getName() string {
	return c.name
}

func (c *MargheritoPizza) getSauce() string {
	return c.sauce
}
func (c *MargheritoPizza) getToppings() []string {
	return c.toppings
}
func (c *MargheritoPizza) setToppings(topping []string) {
	c.toppings = topping
}
