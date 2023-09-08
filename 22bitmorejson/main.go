package main

import (
	"encoding/json"
	"fmt"
)

type cource struct {
	Name      string `json:"courcename"`
	price     int
	plateform string   `json:"website"`
	password  string   `json:"-"`
	tags      []string `json:"tags,omitempty"`
}

func encodeJson() {
	lcoCource := []cource{
		{"reactjs", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Solidity", 479, "LearnCodeOnline.in", "cyrpt677", nil},
		{"MERN", 799, "LearnCodeOnline.in", "react4212", []string{"React", "Mongo", "Node", "js", "Express"}},
	}

	finalJson, err := json.MarshalIndent(lcoCource, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func main() {
	encodeJson()
}
