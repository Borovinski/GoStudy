package main

import "fmt"

func main() {
	name := "Americano"
	price := 2.99
	ready := true
	orderedCount := 5
	var stockCount int64 = 5000

	fmt.Printf("Type of name is: %T\n", name)
	fmt.Printf("Type of name is: %T\n", price)
	fmt.Printf("Type of name is: %T\n", ready)
	fmt.Printf("Type of name is: %T\n", orderedCount)
	fmt.Printf("Type of name is: %T\n", stockCount)
}
