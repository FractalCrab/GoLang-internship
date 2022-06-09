// variable declaration
package main

import (
	"fmt"
)

var i bool = true

var (
	j     int    = 42
	hello string = "Hello"
)

func main() {
	// var i int = 42 // declare but not assign
	// variable with innermost scope takes precedence
	// variables always have to be used
	y := "A"
	fmt.Println(i)
	fmt.Println(y)
}
