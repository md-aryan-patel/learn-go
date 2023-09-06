package main

import "fmt"

// Differ work on LIFO
func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("World")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, ", ")
	}
}
