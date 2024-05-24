package main

import "fmt"

type User struct {
	Name string
}

func main() {
	user := getUser()
	fmt.Println(user.Name)
}

func getUser() (user User) {
	user.Name = "Alice"
	return user
}
