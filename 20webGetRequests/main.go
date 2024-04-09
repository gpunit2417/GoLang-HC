package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Handling get requests in golang")
	performGetRequests()
}

func performGetRequests(){
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)

	if err!= nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code is:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	//method 1:
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is:", byteCount)
	fmt.Println(responseString.String())
	
	//method 2:
	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(content)
	// fmt.Println("The content is:", string(content))
}