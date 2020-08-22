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
}
