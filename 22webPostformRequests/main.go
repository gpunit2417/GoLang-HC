package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Handling web postform requests in golang")
	PerformPostFormRequests()
}

func PerformPostFormRequests(){
	const myurl = "http://localhost:1111/post"

	//form data
	data := url.Values{}
	data.Add("firstName", "Punit")
	data.Add("lastName", "Goyal")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}