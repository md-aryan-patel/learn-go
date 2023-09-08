package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func PerformGetRequest() {
	myurl := createUrl("http", "localhost:8080", "get")
	fmt.Println(myurl)
	response, err := http.Get(myurl)
	handleNilError(err)

	defer response.Body.Close()

	fmt.Println(response)

	fmt.Println("Status: ", response.Status)
	content, err := io.ReadAll(response.Body)
	handleNilError(err)

	// Creating the string builder
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func performPostJsonRequest() {
	myurl := createUrl("http", "localhost:8080", "post")
	fmt.Println(myurl)

	requestBody := strings.NewReader(`
		{
			"cource":"reactjs",
			"price":0,
			"plateform": "LearnCodeOneline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	handleNilError(err)
	defer response.Body.Close()
	content, contentErr := io.ReadAll(response.Body)
	handleNilError(contentErr)
	var responseString strings.Builder
	_, byteErr := responseString.Write(content)
	handleNilError(byteErr)
	fmt.Println(responseString.String())
}

func perfoemPostformJsonRequest() {
	// formdata
	myurl := createUrl("http", "localhost:8080", "postform")
	data := url.Values{}
	data.Add("Firstname", "Captain")
	data.Add("Lastname", "Price")
	data.Add("Email", "pricepaid@cod.com")
	response, err := http.PostForm(myurl, data)
	handleNilError(err)
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	var responseString strings.Builder
	_, byteErr := responseString.Write(content)
	handleNilError(byteErr)
	fmt.Println(responseString.String())
}

func main() {
	// PerformGetRequest()
	// performPostJsonRequest()
	perfoemPostformJsonRequest()
}

/* Helper functions */
func handleNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func createUrl(scheme string, host string, path string) string {
	url := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	return url.String()
}
