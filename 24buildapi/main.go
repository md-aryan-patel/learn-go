package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	return c.CourceName == ""
}

func main() {
	fmt.Println("API learnCodeOnLine")
	r := mux.NewRouter()

	// Seeding
	cources = append(cources, Cource{
		CourceId:    "2",
		CourceName:  "React cource",
		CourcePrice: 200,
		Author:      &Author{Fullname: "jerrysuj", Website: "jerry.learncode.dev"}})
	cources = append(cources, Cource{
		CourceId:    "9",
		CourceName:  "MERN",
		CourcePrice: 599,
		Author:      &Author{Fullname: "jerrysuj", Website: "jerry.learncode.dev"}})
	cources = append(cources, Cource{
		CourceId:    "6",
		CourceName:  "Solidity",
		CourcePrice: 199,
		Author:      &Author{Fullname: "jerrysuj", Website: "jerry.learncode.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCources).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCource).Methods("GET")
	r.HandleFunc("/course", createOneCource).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCource).Methods("DELETE")

	// Listening to port
	log.Fatal(http.ListenAndServe(":8080", r))
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

func getOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one cource")
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	// Loop through the cources and get the course
	for _, course := range cources {
		if course.CourceId == param["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("No cource found with the given ID")
	}
}

func createOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create on cource")
	w.Header().Set("Content-Type", "application/json")

	// No body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Plese send some data")
		return
	}

	// data - {}
	var course Cource
	_ = json.NewDecoder(r.Body).Decode(&course)
	if IsEmpty(&course) {
		json.NewEncoder(w).Encode("Empty json")
		return
	}

	// Generate a unique id, convert that id into string
	// append this new course in the db
	rand.Seed(time.Now().UnixNano())
	course.CourceId = strconv.Itoa(rand.Intn(100))
	cources = append(cources, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create on cource")
	w.Header().Set("Content-Type", "application/json")

	// Grabe id from req
	params := mux.Vars(r)

	var course Cource

	for index, cource := range cources {
		if cource.CourceId == params["id"] {
			cources = append(cources[:index], cources[index+1:]...)
			json.NewDecoder(r.Body).Decode(&course)
			course.CourceId = params["id"]
			cources = append(cources, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Data with given id not found")

}

func deleteOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one cource")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range cources {
		if course.CourceId == params["id"] {
			cources = append(cources[:index], cources[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("No data with given index vaule found")
}
