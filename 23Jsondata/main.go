package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to conversion of data into json format series")
	EncodedJson()
	DecodedJson()
}

type course struct {
	//the json value written after the data types are used because we want to show the name as coursename and other values as indicated while printing the json data

	//the `json:"-"` indicates we don't want to show the password field while printing the json data.
	//the omitempty keyword is used to indicate if the tag value is nil then do not print it in json data


	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodedJson() {
	courses := []course{
		{"React JS", 199, "http://lco.dev.in", "abc123", []string{"web-dev", "js"}},
		{"MERN stack", 199, "http://lco.dev.in", "mern123", []string{"mern-dev"}},
		{"Frontend developer", 199, "http://lco.dev.in", "frontend987", nil},
	}

	//package this data as json data

	//marshatIndent is used to package the input data in json format. it takes three arguments: data, indentation tag and the special tag to print in the next line
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodedJson() {
	jsonData := []byte(`
		{
			"coursename": "React JS",      
			"Price": 199,
			"website": "http://lco.dev.in",
			"tags": ["web-dev","js"]
		}
	`)

	var courses course

	//valid tag to check if the input json is a valid json.
	checkValid := json.Valid(jsonData)

	if checkValid{
		fmt.Println("Json is valid")

		//unmarshal unpackages the json data and print it in the courses variable
		json.Unmarshal(jsonData, &courses)
		fmt.Printf("%#v\n", courses)
	} else{
		fmt.Println("Json data is not valid")
	}

	//some cases where you want to add data to a key value so for that we will use map
	var myCourses map[string]interface{}
	json.Unmarshal(jsonData, &myCourses)
	fmt.Printf("%#v\n", myCourses)

	for k, v := range myCourses{
		fmt.Printf("Key is %v and value is %v and the type is: %T\n", k, v, v)
	}
}
