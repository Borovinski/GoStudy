package main

import (
	"fmt"
	"math/rand"
	"time"
)

var number int = 5

func point21(score *int) string {
	var result string = "Ты не попал ни в какую категорию!"

	if *score == 12 {
		result = "Дюжина"
		return result

	} else if *score == 21 {
		result = "Очко"
		return result
	} else if *score == 50 {
		result = "Полтинник"
		return result
	} else {
		return result
	}
}

func mem(score *int) string {
	var result string = "Ты не петушок"
	if (*score < 6) || (*score > 16) {
		result = "Ты петушок"
		return result
	} else {
		return result
	}

}

func makeThrowPipe() func(choice bool) int {
	score := 0

	return func(flewBy bool) int {
		if flewBy {
			score++
			fmt.Println("Вы пролетели трубу")
		} else {
			score = 0
			fmt.Println("Вы врезались в трубу!")
		}
		return score
	}
}

func runBird(choice bool) {
	throwPipe := makeThrowPipe()
	for i := 0; i < 3; i++ {
		score := throwPipe(choice)
		fmt.Println("Счёт:", score)
	}

}

func playBird(numberPipeEndGame int) {
	var score int = 1
	breakPipe := new(int)

	for {
		*breakPipe = rand.Intn(10)
		fmt.Println("Адрес памяти переменной breakPipe:", &breakPipe)
		fmt.Printf("Подлетаю к %d-oй трубе\n", score)
		if numberPipeEndGame == *breakPipe {
			fmt.Printf("Вы врезались в трубу 🔴\n")
			score = 0
			break
		} else {
			fmt.Printf("Вы пролетели %d-ую трубу ✅\n", score)
			score++
			time.Sleep(2 * time.Second)
		}

		fmt.Println("-----------------------------")

	}

}

func greeting(name string) {
	if name == "" {
		fmt.Println("Вы передали пустое имя :(")
		return
	}

	fmt.Println("Привет, уважаемый", name)
}

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func (u User) Greeting() {
	fmt.Println("Всем привет!")
	fmt.Println("Меня зовут:", u.Name)
	fmt.Println("Мой рейтинг:", u.Rating)
}

/*
	func (u User) ValidRating() int {
		if u.Rating > 0 && u.Rating <= 10 {
			return int(u.Rating)
		} else {
			return 0
		}
	}
*/

func (u *User) ValidRating(delta float64) {
	newRating := u.Rating + delta
	if newRating > 0 && newRating <= 10 {
		u.Rating = newRating
	}
}

func main() {
	// playBird(3)

	user := User{
		Name:        "Серега",
		Age:         12,
		PhoneNumber: "+7 (911) 911",
		IsClose:     true,
		Rating:      5.5,
	}
	fmt.Println(user.Rating)
	user.ValidRating(2.0)
	fmt.Println(user.Rating)

	// fmt.Println(user.ValidRating())
}
