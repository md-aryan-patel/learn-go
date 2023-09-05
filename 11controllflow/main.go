package main

import "fmt"

func main() {
	fmt.Println("Conditionals in GO")
	var result string
	loginCounter := 23

	if loginCounter < 10 {
		result = "Good to go"
	} else if loginCounter > 10 {
		result = "Exceed time limit"
	} else {
		result = "Watch out, last login"
	}

	fmt.Println(result)

	if num := 10; num < 10 {
		fmt.Println("Number is less then 10")
	} else if num > 10 {
		fmt.Println("Number is greater then 10")
	} else {
		fmt.Println("Number is eactly 10")
	}
}
