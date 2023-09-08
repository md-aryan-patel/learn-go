package main

import (
	"fmt"
	"sort"
)

func playSlice() {
	var users = []string{"captain", "price", "jerrysujj", "zeus", "yeahboi"}
	users = append(users, "Aryan")
	fmt.Println(users)
}

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruit list is: %T", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(append(fruitList[1:3]))

	highScore := make([]int, 4)
	highScore[0] = 123
	highScore[1] = 612
	highScore[2] = 431
	highScore[3] = 945

	highScore = append(highScore, 828, 756, 777)
	// fmt.Println(highScore)
	sort.Ints(highScore)
	// fmt.Println(highScore)

	// How to remove values using index valuex
	var cources = []string{"reactjs", "javascript", "swift", "python", "kotlin", "solidity"}
	fmt.Println(cources)
	var index int = 2
	// startpoint(inclusive):endpoint(non-inclusive)
	cources = append(cources[:index], cources[index+1:]...)
	playSlice()
}
