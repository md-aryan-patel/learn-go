package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)
	// Creating reader
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Enter your rating")

	// Error ok syntex
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
}