package main

import "fmt"

func main() {
	fmt.Println("This is pointer example")
	var ptr *int
	fmt.Println(ptr)
	
	myNumber := 23
	var ptrMyNumber *int = &myNumber
	fmt.Println(ptrMyNumber)
}