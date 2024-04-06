package main

import "fmt"

const token = 3000 // public variable

func main() {
	var username string = "punit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4554554663344223443442555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var largeFloat float64 = 255.4554554663344223443442555
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	//default value will be 0
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var name = "punit"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	//another method to create a variable using walrus(:) operator
	myName := "punit"
	fmt.Println(myName)
	fmt.Printf("Variable is of type: %T \n", myName)

	//accessing the constant and the public variable created outside the main function
	fmt.Println(token)
	fmt.Printf("Variable is of type: %T \n", token)	
}
