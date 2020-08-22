package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct
	value.myField = 3

	var pointer *myStruct = &value

	pointer.myField = 8

	// fmt.Println(*pointer.myField) can not use this

	// fmt.Println((*pointer).myField) we can use this code to call pointer

	fmt.Println(pointer.myField) // or dot operator (as default) provide us to access field via pointer

	fmt.Println(value.myField)
}
