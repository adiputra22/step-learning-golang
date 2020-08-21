package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5, "Adi": 62}

	var names []string

	for name := range grades {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s punya grade sebesar %0.2f%% \n", name, grades[name])
	}
}
