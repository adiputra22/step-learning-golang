package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed int
}

func main() {
	var porsche car
	porsche.name = "Porsche type AAA"
	porsche.topSpeed = 325

	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolt"
	bolts.count = 24

	fmt.Println("Bolt description:", bolts.description)
	fmt.Println("Bolt Count:", bolts.count)
}
