package main

import "fmt"

func main() {
	fmt.Println("welcome to array series in golang")

	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[3] = "Grapes"

	fmt.Println("The fruit list is:", fruitlist)	//leaves a space for the element at second index
	fmt.Println("The length of fruit list is:", len(fruitlist))		//output is 4 as length of array

	//another method to declare array
	var anotherlist = [3]string{"Punit", "Atul", "Arman"}
	fmt.Println("The another list is:", anotherlist)
	fmt.Println("The length of another list is:", len(anotherlist))
}