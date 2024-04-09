package main

import "fmt"

func main() {
	fmt.Println("Defers in golang")

	//defers follow LIFO property which is of stack to print the stuff
	//defer automatically put the statement with defer keyword at the end of the function, just before the last closing } bracket

	// fmt.Println("Hello")
	// defer fmt.Println("Punit!")

	// fmt.Println("------------------")
	
	// defer fmt.Println("Punit!")
	// fmt.Println("Hello")

	// fmt.Println("------------------")
	
	defer fmt.Println("Punit!")
	defer fmt.Println("Goyal!")
	defer fmt.Println("From!")
	defer fmt.Println("Hisar!")
	fmt.Println("Hello")

	fmt.Println("------------------")

	print()

}

func print(){
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}