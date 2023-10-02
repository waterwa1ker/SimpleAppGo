package models

import "fmt"

func (user *User) SetName(name string) {
	user.name = name
}

func (user User) GetName() string {
	return user.name
}

func (user *User) SetAge(age uint16) {
	user.age = age
}

func (user User) GetAge() uint16 {
	return user.age
}

func (user *User) SetMoney(money int16) {
	user.money = money
}

func (user User) GetMoney() int16 {
	return user.money
}

func (user *User) SetHappiness(happiness float64) {
	user.happiness = happiness
}

func (user User) GetHappiness() float64 {
	return user.happiness
}

func (user *User) SetAverageGrades(averageGrades float64) {
	user.averageGrades = averageGrades
}

func (user User) GetAverageGrades() float64 {
	return user.averageGrades
}

func (user User) GetAllInfo() string {
	return fmt.Sprintf("My name is %s. I'm %d years old! I have %d$", user.name, user.age, user.money)
}
