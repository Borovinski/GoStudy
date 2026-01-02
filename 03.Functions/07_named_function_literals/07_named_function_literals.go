package main

import "fmt"

func main() {
	const taxRate = 0.10
	var subtotal = 25.00

	var calculateTax = func(amount float64) float64 {
		return amount * taxRate
	}

	tax := calculateTax(subtotal)
	total := subtotal + tax
	fmt.Printf("Total amount to pay: $%.2f\n", total)

}
