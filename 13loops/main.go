package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in GO lang")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Fridya", "Saturday"}

	// Reular for loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// blank is for index value
	for _, v := range days {
		fmt.Println(v)
	}

	value := 1
	for value < 10 {
		if value < 5 {
			break
		}
		value++
	}
}
