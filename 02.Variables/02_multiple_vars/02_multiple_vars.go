package main

import "fmt"

var price = 2.50

func main() {
	// Declare and initialize with var
	var coffeeName string = "Espresso"
	var size string = "Medium"
	// var price float32 = 2.50

	fmt.Println("Medium Espresso price is $3.50")
	fmt.Println(size, coffeeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffeeName, price)

}
