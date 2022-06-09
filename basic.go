package main

import (
	"fmt"
)

func main() {
	var first string
	fmt.Println("Hello World")
	fmt.Scanln(&first)
	fmt.Printf("Hello! %v\n", first)
}
