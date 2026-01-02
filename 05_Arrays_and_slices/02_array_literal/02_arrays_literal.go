package main

import "fmt"

func main() {
	var coffeeTypes = [3]string{"Espresso", "Latte", "Cappuccino"}
	fmt.Println(coffeeTypes)
	fmt.Println("Length of the array:", len(coffeeTypes))

	var value string = coffeeTypes[len(coffeeTypes)-1]
	fmt.Println(value)
	fmt.Println("String length is:", len(coffeeTypes))

}
