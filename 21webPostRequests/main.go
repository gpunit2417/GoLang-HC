package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Handling web post requests in golang")
	PerformPostRequests()
}

func PerformPostRequests(){
	const myurl = "http://localhost:1111/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"courseName" : "Golang",
			"Price" : 0,
			"Platform" : "youtube"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}