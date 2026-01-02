package main

import "fmt"

func main() {
	dessertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}
	fmt.Println("Dessert menu:", dessertMenu)
	slice := dessertMenu[1:3]
	fmt.Println(slice)
}
