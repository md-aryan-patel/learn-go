package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/md-aryan-patel/gomongo/router"
)

func main() {
	fmt.Println("Mongodb api")
	fmt.Println("Server is up and runnign")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}
