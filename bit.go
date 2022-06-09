package main

import (
	"fmt"
)

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 40000000000.
	fmt.Printf("%v\n", KB)
	fmt.Printf("%.2fGB \n", fileSize/GB)

}
