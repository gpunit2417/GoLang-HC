package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["jv"] = "java"
	languages["rctj"] = "reactjs"
	languages["cpp"] = "cplusplus"

	fmt.Println("The list of languages: ", languages)
	fmt.Println("Language for js", languages["js"])
	
	delete(languages, "jv")
	fmt.Println("The list of languages: ", languages)

	//loops in golang
	for key, value := range languages{
		fmt.Printf("For key %v, the corresponding language is %v\n", key, value)
	}
	
}