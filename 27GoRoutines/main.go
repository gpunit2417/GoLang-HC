package main

import (
	"fmt"
	"time"
)

func main() {

	//go routines are used to manage the parallelism task means perfoming many task at the same time.
	//it is of 2kb flexible stack size
	//it spawns many threads at a faster pace but is comparable with the threads.
	//it is managed by go runtime.
	fmt.Println("GoRoutines in golang")
	go greeter("Hello")
	greeter("World")
}

// func main() {
// 	fmt.Println("GoRoutines in golang")
// 	greeter("Hello")
// 	greeter("World")
// }

func greeter(s string){
	for i:=0; i<6; i++{

		//without using the below statement we will get onlu world 5 times in the output.
		//this is because the thread is not waiting for the go keyword to get executed with hello argument.
		time.Sleep(3 * time.Millisecond)
		//using the above statement, the thread is waiting for 3ms to get executed with the hello argument.
		//so the order is not fixed in the output.
		fmt.Println(s)
	}
}