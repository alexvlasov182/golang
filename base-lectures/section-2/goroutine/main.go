package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// actions defines a list of possible users actions.
var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

// logItem represents alt entry with action and timestamp
type logItem struct {
	action    string    // action performed by the user
	timestamp time.Time // timestamp of the action
}

// User represent a user wiht an ID, email, and activity log.
type User struct {
	id    int       // unique identifier for the user
	email string    // email addrss of the user
	logs  []logItem // activity log of the user
}

// getActivityInfo returns a string containing users's information and activity log.
func (u User) getActivityInfo() string {
	out := fmt.Sprintf("Id: %d | Email: %s\nActivity Log: \n", u.id, u.email)

	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.timestamp)
	}

	return out
}

// main is the entry point of the program.
// The main function calculates and prints the time elapsed during execution.
func main() {
	rand.Seed(time.Now().Unix())

	startTime := time.Now()

	wg := &sync.WaitGroup{}

	users := make(chan User)

	go generateUsers(1000, users)

	for user := range users {
		wg.Add(1)
		go saveUserInfo(user, wg)
	}

	wg.Wait()

	fmt.Println("Time Elapsed", time.Since(startTime).String())
}

// saveUserInfo saves user information to a file.
func saveUserInfo(user User, wg *sync.WaitGroup) error {
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("Writing file for user id: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}

	wg.Done()

	return nil
}

// generateUsers generates a specified number of users with random logs.
func generateUsers(count int, users chan User) {

	for i := 0; i < count; i++ {
		users <- User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@ninja.go", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
		//time.Sleep(time.Millisecond * 10)
	}

	close(users)

}

// generateLogs generates a specified number of random log entries.
func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			timestamp: time.Now(),
			action:    actions[rand.Intn(len(actions)-1)],
		}
	}

	return logs
}
