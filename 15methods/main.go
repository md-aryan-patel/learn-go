package main

import (
	"fmt"
)

/* if name starts with Capital letter is public, but if
name start with small letter then its not exportable */

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User status is: ", u.Status)
}

func main() {
	newUser := User{"Aryan", "xyz@abc.com", false, 12}
	// way to calling methods
	newUser.GetStatus()
}
