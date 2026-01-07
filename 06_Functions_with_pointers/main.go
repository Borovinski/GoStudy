package main

import "fmt"

func throwPipe(score *int) int {
	fmt.Println("Ваш счет:", *score)
	fmt.Println("Пролетели через трубу")
	*score++
	fmt.Println("Ваш счет:", *score)
	goodScore(score)
	return *score
}

func goodScore(score *int) {
	if *score > 10 {
		fmt.Println("Нормальный результат!")
		if *score > 15 {
			fmt.Println("Лучший результат!")
		} else {
			fmt.Println("Хороший результат")
		}
	} else {
		fmt.Println("Плохой результат!")
	}
}

func breakThrow(score *int) int {
	fmt.Println("Врезались в трубу")
	*score = 0
	return *score
}

func main() {
	const (
		textReady string = "Get Ready"
		gameOver  string = "Game Over"
	)
	var score int = 15
	x := throwPipe(&score)
	fmt.Println(x)

}
