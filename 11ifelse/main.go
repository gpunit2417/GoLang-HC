package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 23
	var result string

	//different ways of if else statement

	if loginCount < 10{
		result = "login count less than 10"
	}else if loginCount > 10{
		result = "login count more than 10"
	}else{
		result = "something else"
	}

	fmt.Println(result)


	if 9%2 == 0{
		fmt.Println("Number is even")
	} else{
		fmt.Println("Number is odd")
	}


	if num := 3; num < 10{
		fmt.Println("Number less than 10")
	} else{
		fmt.Println("Number more than or equal to 10")
	}
}