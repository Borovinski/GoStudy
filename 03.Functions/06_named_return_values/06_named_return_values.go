package main

import "fmt"

// func estimateBrewTime(cupsQty int, secondsPerCup int) int {
// 	totalTimeSeconds := cupsQty * secondsPerCup
// 	return totalTimeSeconds

// }

func estimateBrewTime(cupsQty int, secondsPerCup int) (totalTimeSeconds int, info string) {
	totalTimeSeconds = cupsQty * secondsPerCup
	info = "Estimated:"
	return

}
func main() {
	breTime, info := estimateBrewTime(12, 20)
	fmt.Println(info, breTime)
}
