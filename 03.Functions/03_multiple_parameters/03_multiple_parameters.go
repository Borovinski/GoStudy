package main

import "fmt"

func getDrinkInfo(customerName string, drink string, price float64) {
	info := "%s favorite drink is %s %.2f\n"
	fmt.Printf(info, customerName, drink, price)
}

func main() {
	getDrinkInfo("Dima", "Coffee", 4.50)
}
