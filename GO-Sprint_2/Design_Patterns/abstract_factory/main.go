// main.go
package main

import "fmt"

func PrintShoeDetails(shoe IShoe) {
	fmt.Printf("The shoe logo is %s\n", shoe.getLogo())
	fmt.Printf("The shoe size is %d\n", shoe.getSize())
}

func PrintShirtDetails(shirt IShirt) {
	fmt.Printf("The shirt logo is %s\n", shirt.getLogo())
	fmt.Printf("The shirt size is %d\n", shirt.getSize())
}

func main() {
	nikeFactory, _ := GetSportFactory("nike")
	adidasFactory, _ := GetSportFactory("adidas")

	nikeShirt := nikeFactory.makeShirt()
	nikeShoe := nikeFactory.makeShoe()

	adidasShirt := adidasFactory.makeShirt()
	adidasShoe := adidasFactory.makeShoe()

	PrintShirtDetails(nikeShirt)
	PrintShoeDetails(nikeShoe)

	PrintShirtDetails(adidasShirt)
	PrintShoeDetails(adidasShoe)
}
