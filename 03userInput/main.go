package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the user input:")

	//comma ok || err ok syntax to handle try and catch like functioning

	// input, err := reader.ReadString('\n')
	// _, err := reader.ReadString('\n')

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks to enter: ", input);
	fmt.Printf("The type of the user input is: %T", input)
}