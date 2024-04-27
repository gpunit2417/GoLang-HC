package main

import (
	"fmt"
	"sync"
)

//channels are a way of communication between the go routines 
//without letting them aware of what is going on inside them
func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 1)	//creating a channel

	// myCh <- 5
	// fmt.Println(myCh)  we cannot print like this as this is giving error
	wg := &sync.WaitGroup{}

	wg.Add(2)

	//Method 1: -

	//recieve
	// go func (ch <-chan int, wg *sync.WaitGroup)  {
	// 	val, isChannelOpen := <-myCh	//to check if there is anything present inside the channel or not

	// 	fmt.Println(isChannelOpen)	//return true if present else false
	// 	fmt.Println(val)	//return the value inside the channel

	// 	wg.Done()
	// }(myCh, wg)

	// //send 
	// go func (ch chan<- int, wg *sync.WaitGroup)  {
	// 	myCh <- 0
	// 	close(myCh) 	//closing the channel

	// 	//myCh <- 6
	// 	wg.Done()
	// }(myCh, wg)

	//Method 2: -
	//recieve
	go func (ch chan int, wg *sync.WaitGroup)  {
		
		fmt.Println(<-myCh)	
		fmt.Println(<-myCh)	

		wg.Done()
	}(myCh, wg)

	//send 
	go func (ch chan int, wg *sync.WaitGroup)  {
		myCh <- 5
		myCh <- 6

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}