package main

import (
	"fmt"
	"time"
)

func main() {
	welcome := "welcome to handling time series"
	fmt.Println(welcome)

	//time package is used to get the time and Now() method gives the current time of your system.
	presentTime := time.Now()
	fmt.Println(presentTime)

	//to format the output of the previous statement, we use Format() method which takes the below argument.
	//NOTE:- The argument are fixed. Don't change these. otherwise you will get undesired output.
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//Date() method of time package is used to create a new method
	//this method takes the below 8 parameters viz, year, month, day, hour, min, sec, nanosec, location
	createdDate := time.Date(2024, time.April, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}