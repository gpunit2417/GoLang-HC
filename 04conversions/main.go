package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the user rating:")

	//comma ok || err ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for entering:", input)

	//adding 1 to the existing rating

	//numRating is another variable to store the rating that user give
	//strconv is a golang package used to convert string to int type
	//ParseFloat is used to convert the datatype to float datatype. It takes 2 arguments: variable to convert and the bit of conversion
	//strings is a package in golang and TrimSpace is a method of this package to trim the given input to only number as input here is "4\n"
	
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Added 1 to the rating:", numRating+1)
	}
}