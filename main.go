package main

import (
	"fmt"

	"github.com/adiputra22/magazine-go"
)

func main() {
	address := magazine.Address{Street: "Jl Sabre No8A Margahayu", City: "Bekasi Timur", State: "Jawa Barat", PostalCode: "17113"}

	subscriber1 := magazine.Subscriber{Name: "Adiputra", HomeAddress: address}

	subscriber1.Active = true
	subscriber1.Rate = 800

	fmt.Println(subscriber1)

	subscriber1.HomeAddress.Street = "Jl Ampera No64 Duren Jaya"

	fmt.Println(subscriber1.HomeAddress.Street)

	employee := magazine.Employee{Name: "Irsyad", Salary: 5000}

	employee.HomeAddress.Street = "Jl Perwira 2"
	employee.HomeAddress.City = "Bekasi"
	employee.HomeAddress.State = "Jawa Barat"
	employee.HomeAddress.PostalCode = "17123"

	fmt.Println(employee)
}
