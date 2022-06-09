package main

import "fmt"

const (
	catSpecialist = iota + 5
	dogSpecialist
	snakeSpecialist
)

func main() {
	// var specialListType int
	// fmt.Printf("%v \n", dogSpecialist)
	// fmt.Printf("%v \n", snakeSpecialist)
	// fmt.Printf("%v\n", specialListType == dogSpecialist)
	fmt.Printf("%v\n", catSpecialist)
}
