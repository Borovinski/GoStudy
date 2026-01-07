package main

import "fmt"

type User struct {
	//Имя пользователя не должно быть пустое
	Name string
	//Возраст не меньше нуля и не больше 150
	Age int
	//Не должно быть пустым
	PhoneNumber string
	//Закрыт ли профиль пользователя
	IsClose bool
	//Рейтинг от 0.0-10.0
	Rating float64
}

func NewUser(name string, age int, phoneNumber string, isClose bool, rating float64) (User, error) {
	if name == "" {
		return User{}, fmt.Errorf("имя не может быть пустым")
	}
	if age <= 0 || age > 150 {
		return User{}, fmt.Errorf("возраст должен быть от 1 до 150")
	}
	if phoneNumber == "" {
		return User{}, fmt.Errorf("номер телефона не может быть пустым")
	}
	if rating < 0.0 || rating > 10.0 {
		return User{}, fmt.Errorf("рейтинг должен быть от 0.0 до 10.0")
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
	}, nil
}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	} else {
		fmt.Println("У вас пустое имя!")
	}
}

func (u *User) ChangeAge(newAge int) {
	validAge := newAge > 0 && newAge < 150
	if validAge {
		u.Age = newAge
	} else {
		fmt.Println("Ваш возраст не определен")
	}

}

func (u *User) ChangePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	} else {
		fmt.Println("У вас пустой номер!")
	}
}

func (u *User) CloseAccount(choice bool) {
	if choice {
		u.IsClose = true
	} else {
		u.IsClose = false
	}
}

func (u *User) ChangeRating(newRating float64) {
	validRating := newRating >= 0 && newRating <= 10
	if validRating {
		u.Rating = newRating
	} else {
		fmt.Println("Невалидный рейтинг!")
	}
}

func main() {
	// var user1 User = User{
	// 	Name:        "Дмитрий",
	// 	Age:         26,
	// 	PhoneNumber: "+375339926673",
	// 	IsClose:     false,
	// 	Rating:      2.0,
	// }

	user, error := NewUser(
		"Данил", 12, "3434", true, 2.0,
	)
	fmt.Println(user, error)

	// fmt.Println(user1)
	// user1.ChangeName("Слава")
	// user1.ChangeAge(65)
	// user1.ChangePhoneNumber("345")
	// user1.CloseAccount(true)
	// user1.ChangeRating(3.0)
	// fmt.Println((user1))
}
