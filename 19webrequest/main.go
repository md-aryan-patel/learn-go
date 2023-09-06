package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

type User struct {
}

func main() {
	res, err := http.Get(url)
	checkNilError(err)
	fmt.Printf("Response is of type: %T", res)
	defer res.Body.Close() // It is callers responsibility to close the connection
	databytes, err := io.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println("Data from response: ", string(databytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
