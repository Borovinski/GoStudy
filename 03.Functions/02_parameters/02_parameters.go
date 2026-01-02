package main

import "fmt"

func greet(coffeeShopName string) {
	// var coffeeShopName string = "Down"
	fmt.Println("Welcome to the Coffee Shop", coffeeShopName)

}

func main() {
	greet("Blue&Beans")
}
