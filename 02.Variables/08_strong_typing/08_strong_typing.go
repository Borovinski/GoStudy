package main

import "fmt"

func main() {
	//price of one cup of coffee
	price := 4.50

	//cups sold in one day
	quantity := 15

	//total income
	// total := price * quantity
	// fmt.Println(total)

	total := int(price) * quantity
	fmt.Printf("Total income during a day %d\n", total)

}
