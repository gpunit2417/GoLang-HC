package main

import "fmt"

func main() {
	fmt.Println("loops, break, continue and goto statement in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "friday", "Saturday", "Sunday"}
	fmt.Println(days)

	for d := 0; d<len(days); d++{
		fmt.Println("The days are: ", days[d])
	}

	for i := range days{
		fmt.Println(days[i])
	}

	for index, day := range days{
		fmt.Printf("The index is %v and the day is %v", index, day)
	}


	value := 1

	for value<10{
		fmt.Println("Value is:", value)
		value++
	}

	for value<10{
		if value == 5{
			break
		}
		fmt.Println("Value is:", value)
		value++
	}

	for value<10{
		if value == 5{
			value++
			continue
		}
		fmt.Println("Value is:", value)
		value++
	}

	if value == 1{
		goto punit
	}

	punit:
	fmt.Println("This is a goto syntax")
}