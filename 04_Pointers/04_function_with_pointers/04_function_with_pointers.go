package main

import "fmt"

func applyDiscount(price *float64, discountRate *float64) {
	fmt.Println("Memory location of the price in calculatePriceAfterDiscount:", price)
	*price = *price - *price**discountRate

}

func main() {
	var coffePrice float64 = 5.00
	fmt.Println("Mem location coffePrice in main:", &coffePrice)
	var discount float64 = 0.10
	applyDiscount(&coffePrice, &discount)
	fmt.Println("Basic coffee price:", coffePrice)
}
