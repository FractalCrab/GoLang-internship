package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Himachal": 70,
		"UP":       220,
	}
	// m := map[[3]int]string{}
	// fmt.Println(m)
	statePopulations["UK"] = 50
	delete(statePopulations, "UK")
	pop, ok := statePopulations["Himachal"]
	fmt.Println(pop, ok)
}
