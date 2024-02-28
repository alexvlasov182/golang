package main

import "fmt"

func main() {
	// Maps does not have capacity
	users := map[string]int{
		"John":  14,
		"Jack":  33,
		"Jemmy": 43,
	}

	users["Jill"] = 20
	delete(users, "John")
	fmt.Println(users)
	fmt.Println(users["John"])
	fmt.Println("Length of Users: ", len(users))

	var people map[string]int
	people = make(map[string]int)
	people["Keanu"] = 23

	fmt.Println(people)
	fmt.Println("Length of People: ", len(people))

	age, exists := users["Jack"]
	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("This user not in the list")
	}

	for name, age := range users {
		fmt.Println("Name: ", name, " - age: ", age)
	}
}
