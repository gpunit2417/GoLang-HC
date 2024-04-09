package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.flipkart.com/viewcart?exploreMode=true&preference=FLIPKART"

func main() {
	fmt.Println("Handling URLs in golang")
	fmt.Println(myurl)
	
	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is: %T\n", qparams)

	fmt.Println(qparams["exploreMode"])

	for _, val := range qparams{
		fmt.Println("Param is:", val)
	}

	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "www.flipkart.com",
		Path: "viewcart",
		RawQuery: "exploreMode=true&preference=FLIPKART",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println(anotherUrl)
}