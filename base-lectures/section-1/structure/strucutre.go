package main

import "fmt"

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

// Pointer receiver
func (u *User) setNmae(name string) {
	u.name = name
}

// Value receiver
func (u User) getName() string {
	return u.name
}

type DumbDatabase struct {
	m map[string]string
}

func newDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    age,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Galex", "Male", 23, 100, 188)
	user2 := User{"Olek", 36, "Male", 90, 188}

	user1.setNmae("Serega")
	fmt.Println(user1)
	fmt.Println(user1.getName())
	fmt.Println(user2.getName())
}
