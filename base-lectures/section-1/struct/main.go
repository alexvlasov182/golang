// Main Package
package main

import (
	"fmt"
)

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func (u *User) setName(name string) {
	u.name = name
}

func (u User) getName() string {
	return u.name
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("John", "Male", 23, 75, 188)
	user2 := User{"Carl", 43, "Male", 85, 199}

	user1.setName("Serega")
	fmt.Println(user1.getName())
	fmt.Println(user2.getName())

	fmt.Println(user1.age)
}
