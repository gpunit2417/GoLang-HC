package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to conversion of data into json format series")
	EncodedJson()
}

type course struct{
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func EncodedJson(){
	courses := []course{
		{"React JS", 199, "http://lco.dev.in", "abc123", []string{"web-dev", "js"}},
		{"MERN stack", 199, "http://lco.dev.in", "mern123", []string{"mern-dev"}},
		{"Frontend developer", 199, "http://lco.dev.in", "frontend987", nil},
	}

	//package this data as json data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err!= nil{
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}