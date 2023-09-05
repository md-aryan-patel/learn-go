package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["sol"] = "solidity"
	languages["py"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	delete(languages, "py")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println("key is: ", key, "value is: ", value)
	}
}
