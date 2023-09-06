package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?courcename=reactjs&paymentid=shfah1234"

func createUrl() {
	url := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=ryan",
	}

	result := url.String()
	fmt.Println(result)
}

func destructureUrl() {
	fmt.Println("This is URL")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Raw query: ", result.RawQuery)
	fmt.Println("Port: ", result.Port())

	qparams := result.Query()

	for k, v := range qparams {
		fmt.Println(k, ":", v)
	}
}

func main() {
	createUrl()
}
