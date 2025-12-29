package main

import "fmt"

func main() {
	//нетипизированная константа
	const RewardPoints = 10

	fmt.Printf("Default type of RewardPoints is %T\n", RewardPoints)
	var totalRewardPoints float64 = 150.3
	totalRewardPoints += RewardPoints
	fmt.Printf("Updated points %.2f\n", totalRewardPoints)

}
