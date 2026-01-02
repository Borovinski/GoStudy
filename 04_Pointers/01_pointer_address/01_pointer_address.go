package main

import "fmt"

func main() {
	coffee := "Espresso"
	pointer := &coffee

	fmt.Println("Coffe name coffee:", coffee)
	fmt.Println("Memory location:", pointer)
	fmt.Printf("Pointer address: %p\n", pointer)

	coffeeCopy := coffee

	fmt.Println("Coffe name coffecopy:", coffeeCopy)
	fmt.Println("Memory location:", &coffeeCopy)
	fmt.Printf("Pointer address: %p\n", &coffeeCopy)

}
