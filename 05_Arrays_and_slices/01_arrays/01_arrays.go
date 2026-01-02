package main

import "fmt"

func main() {
	var coffeSizes [3]string
	coffeSizes[0] = "Small"
	coffeSizes[1] = "Medium"
	coffeSizes[2] = "Large"
	fmt.Println("---------------------------------------")
	fmt.Println(len(coffeSizes))
	fmt.Println(coffeSizes)
}
