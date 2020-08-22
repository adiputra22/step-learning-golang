package main

import "fmt"

// example MyType
type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (m MyType) sayHiWithParameter(number int, flag bool) {
	fmt.Println(m)

	fmt.Println(number)

	fmt.Println(flag)
}

func (m MyType) sayHiWithReturn(number int) int {
	return number * 2
}

// type MyNumber
type MyNumber int

// need to add pointer if you want that number changed!
func (number *MyNumber) makeDouble() {
	*number *= 2
}

func main() {
	value := MyType("adiputra")
	value.sayHi()

	anotherValue := MyType("bude")
	anotherValue.sayHi()

	otherValue := MyType("Other")
	otherValue.sayHiWithParameter(4, true)

	returned := otherValue.sayHiWithReturn(2)

	fmt.Println("Returned value:", returned)

	number := MyNumber(2)
	fmt.Println("Original number is:", number)

	number.makeDouble()
	fmt.Println("Angka berubah menjadi:", number)
}
