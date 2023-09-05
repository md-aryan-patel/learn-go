package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	age    int
}

func main() {
	// No inheritace in GO, No super/Parent in GO lang
	fmt.Println("Structs in GO lang")

	// creating user
	newUser := User{"Aryan", "xyz@abc.com", false, 12}
	fmt.Printf("%+v", newUser)
	fmt.Println("User name is ", newUser.Name)
}
