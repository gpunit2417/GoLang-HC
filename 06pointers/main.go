package main

import "fmt"

func main() {
	//declaring a variable which is pointing to null
	// var ptr *int
	// fmt.Println("The value of *int is: ", ptr)


	myNumber := 23
	var ptr = &myNumber	//reference to a variable 

	fmt.Println("The value of ptr is:", ptr)	//address of a variable
	fmt.Println("The value of *ptr is:", *ptr)	//value at the address

	*ptr = *ptr * 2;
	fmt.Println("The value of *ptr after modification: ", *ptr)
	fmt.Println("The value of ptr after modification: ", ptr)
}