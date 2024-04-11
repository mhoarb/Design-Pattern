package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()
	printShoeDetails(nikeShoe)
	printShoeDetails(nikeShirt)
	printShoeDetails(adidasShoe)
	printShoeDetails(adidasShirt)

}

func printShoeDetails(s IShoe) {
	fmt.Println(s.getLogo(), s.getSize())
}
