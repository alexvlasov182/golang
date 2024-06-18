package main

import "fmt"



type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Sex    string `json:"sex"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
}

// Pointer receiver
func (u *User) setNmae(name string) {
	u.Name = name
   
}

// Value receiver
func (u User) getName() string {
	return u.Name
}

type DumbDatabase struct {
	m map[string]string
}

func newDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func NewUser(name, sex string, age, weight, height int) User  {
	return User{
		Name:   name,
		Sex:    sex,
		Age:    age,
		Weight: weight,
		Height: height,
	}
}

func main() {
	user1 := NewUser("Galex", "Male", 23, 100, 188)
	user2 := User{"Olek", 36, "Male", 90, 188}

	user1.setNmae("Serega")
	fmt.Println(user1)
	fmt.Println(user1.getName())
	fmt.Println(user2.getName())
	fmt.Println(newDumbDatabase())
}
