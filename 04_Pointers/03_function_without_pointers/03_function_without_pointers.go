package main

import "fmt"

func calculatePriceAfterDiscount(price float64, discountRate float64) float64 {
	var newPrice = price - price*discountRate
	return newPrice
}

func main() {
	var coffePrice float64 = 5.00
	var discount float64 = 0.10
	coffePrice = calculatePriceAfterDiscount(coffePrice, discount)
	fmt.Println("Basic coffee price:", coffePrice)
}
