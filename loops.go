package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

}
