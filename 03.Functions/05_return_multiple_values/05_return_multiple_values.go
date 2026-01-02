package main

import "fmt"

func processPayment(orderTotal float64, tip float64, amountPaid float64) (float64, float64) {
	totalAmountDue := orderTotal + tip
	change := amountPaid - totalAmountDue
	return totalAmountDue, change
}

func main() {
	totalCost, change := processPayment(6.50, 2.00, 10.00)
	fmt.Println(totalCost, change)
}
