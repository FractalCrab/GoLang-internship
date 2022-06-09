package main

import (
	"fmt"
)

// Main function
func main() {
	grades := [...]int{88, 77, 66}
	// grades[4] = 99
	var students [3]string
	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Students: %v", students)
}
