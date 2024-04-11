package main

type Adidas struct {
}

/*
*
a method for creating adidas shoe
*/
func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

/*
*
a method for creating adidas shirt that return an interface (iShirt)
*/
func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 10,
		},
	}
}
