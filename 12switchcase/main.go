package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The random number generated is:", diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("You can open the game by moving 1 spot")
	
	case 2:
		fmt.Println("Move 2 spots ahead")

	case 3:
		fmt.Println("Move 3 spots ahead")

	case 4:
		fmt.Println("Move 4 spots ahead")

	case 5:
		fmt.Println("Move 5 spots ahead")

	case 6:
		fmt.Println("Move 6 spots ahead")

	default:
		fmt.Println("Invalid dice number")
	}
}