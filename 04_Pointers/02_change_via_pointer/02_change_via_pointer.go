package main

import "fmt"

//Если мы присваиваем к одной и той же переменной новое значение, то адресс памяти всегда будет одинаковым!
func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffe price:", coffeePrice)
	// pointer := &coffeePrice
	fmt.Println("Memory address:", &coffeePrice)

	coffeePrice = 5.00

	fmt.Println("Coffe price:", coffeePrice)
	fmt.Println("Memory address:", &coffeePrice)

	// pointerToCoffePrice := &coffeePrice

	var pointerToCoffePrice *float64 = &coffeePrice
	*pointerToCoffePrice = 5.50
	fmt.Println(coffeePrice)

}
