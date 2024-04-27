package main

var _ Pizza = &VeggieMania{}

type VeggieMania struct {
	price int
}

func (v *VeggieMania) getPrice() int {
	return v.price
}
