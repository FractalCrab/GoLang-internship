package main

import "fmt"

//one return
func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

//multiple return values
func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

func main() {
	x := 5.75
	y := 6.25

	result := avg(x, y)
	stock, _ := getStockPriceChange(100, 200)

	fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result)
	fmt.Printf("Stock price = %.2f", stock)
}
