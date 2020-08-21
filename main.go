package main

import "fmt"

func main() {
	status("Alma")
	status("Adi")
}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}

	grade, ok := grades[name]

	if !ok {
		fmt.Printf("nama %s tidak ada di record!\n", name)
	} else {
		if grade < 60 {
			fmt.Printf("%s is failing!\n", name)
		}
	}
}
