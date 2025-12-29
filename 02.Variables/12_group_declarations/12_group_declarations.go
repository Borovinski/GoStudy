package main

import "fmt"

func main() {
	var coffeType string = "Latte"
	var quantity int = 3
	var unitPtice float64 = 4.25

	fmt.Printf("Ordered %d %s prices %.2f\n", quantity, coffeType, unitPtice)

	var (
		customerName string = "Bogdan"
		tableNumber  int    = 8
		isReadyToPay bool   = false
	)

	fmt.Printf("Customer %s at table %d is ready to pay: %t\n", customerName, tableNumber, isReadyToPay)

	const (
		SizeSmall   = "S"
		SizedMedium = "M"
		SizeLarge   = "L"
	)
}
