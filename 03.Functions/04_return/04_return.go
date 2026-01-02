package main

import "fmt"

func updateTotalPoints(currentPoints int, newPoints int) int {
	return currentPoints + newPoints
}

func calculateLoyaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2)
	return loyaltyPoints
}

func main() {
	var totalPoints int = 120
	newTotalPoints := calculateLoyaltyPoints(9.50)
	fmt.Println(updateTotalPoints(totalPoints, newTotalPoints))
}
