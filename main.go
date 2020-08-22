package main

import "fmt"

type part struct {
	description string
	count       int
}

func main() {
	var bolts part

	bolts.description = "ini bolts description"
	bolts.count = 1243

	showInfo(bolts)
}

func showInfo(p part) {
	fmt.Println("Description bolts:", p.description)
	fmt.Println("Count bolts:", p.count)

	newPart := setPartByDescription("bolt 123")
	fmt.Println("Description new part:", newPart.description)
	fmt.Println("Count new part:", newPart.count)
}

func setPartByDescription(description string) part {
	var p part
	p.description = description
	p.count = 5

	return p
}
