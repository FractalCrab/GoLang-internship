package main

import (
	"fmt"
)

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}
func main() {
	area, _ := rectProps(10.2, 5.4)
	fmt.Printf("area = %f ", area)
}
