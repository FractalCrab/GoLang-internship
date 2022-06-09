//strings
package main

import (
	"fmt"
)

func main() {
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T \n", s, s)
	fmt.Printf("%v, %T \n", b, b)
}
