package main

import "fmt"

func greater() {
	fmt.Println("Hello world!")
}

func addNumber(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func main() {
	fmt.Println("This is function example")
	greater()
	// result := addNumber(12, 4)
	proResult := proAdder(12, 332, 434, 123, 99)
	fmt.Println(proResult)
}
