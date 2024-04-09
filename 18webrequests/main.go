package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://react.dev/reference/react-dom/components#all-svg-components"

func main() {
	fmt.Println("Handling web requests in golang")

	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)	

	defer response.Body.Close()	//necessary to close the response connection

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil{
		panic(err)
	}

	content := string(databyte)
	fmt.Println("The content of the url is:", content)
}