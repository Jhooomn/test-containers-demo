package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Users []User

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (u User) ToString() string {
	return fmt.Sprintf("User ID: %d, Name: %s, Email: %s, Username: %s", u.ID, u.Name, u.Email, u.Username)
}

func getRandomName() string {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Hannah", "Ivan", "Julia"}
	return names[rand.Intn(len(names))]
}

func getRandomEmail(name string) string {
	domains := []string{"example.com", "mail.com", "test.org"}
	return fmt.Sprintf("%s@%s", name, domains[rand.Intn(len(domains))])
}

func getRandomUsername(name string) string {
	num := rand.Intn(100)
	return fmt.Sprintf("%s%d", name, num)
}

func GetUsers() Users {
	users := make([]User, 10)
	for i := 0; i < 10; i++ {
		name := getRandomName()
		users[i] = User{
			ID:       i + 1,
			Name:     name,
			Email:    getRandomEmail(name),
			Username: getRandomUsername(name),
		}
	}
	return users
}

func (u Users) GetRandomUser() *User {
	if len(u) == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(u))
	return &u[randomIndex]
}
