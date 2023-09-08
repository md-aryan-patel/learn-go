package main

import (
	"fmt"
	"io"
	"os"
)

func createFile(filename string, content string) {
	// Creating file in directory
	file, err := os.Create("./" + filename)
	// Handling Error
	if err != nil {
		panic(err)
	}
	// Creating Writetring buffer
	len, err := io.WriteString(file, content)
	// Handling Error
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Length is: ", len)
	}
	// Closing file
	defer file.Close()
}

func readFile(filename string) {
	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("File data: ", string(dataBytes))
}

func main() {
	fmt.Println("Files system in GO lang")
	fileName := "go-fs-test.txt"
	// content := "This is file content, do not erase this"
	// Creating the file
	// createFile(fileName, content)
	readFile(fileName)
}
