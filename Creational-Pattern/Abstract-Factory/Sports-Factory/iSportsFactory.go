package main

import "fmt"

/*
*
abstract product interface
*/
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {

		return &Nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type")
}
