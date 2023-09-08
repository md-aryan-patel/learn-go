package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Plese enter you rating: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// Parsing to flot
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// Parsing to int
	intrating, _ := strconv.ParseInt(input, 10, 64)
	fmt.Println(intrating)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to rating, ", numRating+1)
	}
}
