package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toogle bool
	}

	myStruct.number = 25.2
	myStruct.word = "hai struct"
	myStruct.toogle = true

	fmt.Printf("%#v\n", myStruct)

	fmt.Println(myStruct.number)
}
