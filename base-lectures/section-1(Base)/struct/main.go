// Main Package
package main

import "fmt"

type User struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Sex    string `json:"sex,omitempty"`
	Weight int    `json:"weight,omitempty"`
	Height int    `json:"height,omitempty"`
}

func (u User) getName() string {
	return u.Name
}

type DumbDatabase struct {
	m map[string]string
}

// NewDumbDatabase function  
func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		Name:   name,
		Age:    age,
		Sex:    sex,
		Weight: weight,
		Height: height,
	}
}

func main() {
	user1 := NewUser("John", "Male", 23, 75, 188)
	user2 := NewUser("Jeremy", "Male", 19, 45, 156)

	fmt.Println("Hello Golang")

	fmt.Println(user1)
	fmt.Println(user2)
}
