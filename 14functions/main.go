package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")

	greet()

	result := adder(3, 5)
	fmt.Println("Adder function called and the result of 3+5 is:", result)

	proResult := autoAdder(3, 4, 5, 6)
	fmt.Println("From the autoAdder function. the result is:", proResult)

	proResult1, myMessage := autoAdder1(3, 4, 5, 6, 7)
	fmt.Println("From the autoAdder1 function. the result is:", proResult1)
	fmt.Println("From the autoAdder1 function. the message is:", myMessage)
}

func greet(){
	fmt.Println("Greeting from the greet function")
}

func adder(varOne int, varTwo int) int{
	return varOne + varTwo
}

func autoAdder(values ...int) int{
	total := 0

	for _, val := range values{
		total += val
	}

	return total
}

func autoAdder1(values ...int) (int, string){
	total := 0

	for _, val := range values{
		total += val
	}

	return total, "Hi there! passing a message with the returned value"
}