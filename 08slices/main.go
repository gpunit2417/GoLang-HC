package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slicing series on array")

	var fruitlist = []string{"Apple", "Banana", "Peach"}
	fmt.Printf("The type of fruitlist is: %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Grapes")
	fmt.Println("The fruitlist after modification:")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println("Slicing the fruitlist array:")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println("Slicing the fruitlist array:")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:])
	fmt.Println("Slicing the fruitlist array:")
	fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0] = 243
	highScores[1] = 944
	highScores[2] = 543
	highScores[3] = 843
	//highScores[4] = 943	 //adding elements like this will give error array out of bound

	fmt.Println(sort.IntsAreSorted(highScores))
	highScores = append(highScores, 55, 666, 897)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}