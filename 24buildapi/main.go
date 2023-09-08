package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model - file/folder
type Cource struct {
	CourceId    string  `json:"courceid"`
	CourceName  string  `json:"courcename"`
	CourcePrice int     `json:price`
	Author      *Author `json:author`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var cources []Cource

// middleware
func IsEmpty(c *Cource) bool {
	return c.CourceId == "" && c.CourceName == ""
}

func main() {

}

// Controllers - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Learncodeonline</h1>"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all cources")
	// setting header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources)
}

func getCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get cource")
	courceId := mux.Vars(r)["courceId"]
}
