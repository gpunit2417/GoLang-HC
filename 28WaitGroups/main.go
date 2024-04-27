package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup	//pointer variable of a wait group object wg
var mut sync.Mutex  //pointer variable of a mutex object mut

func main() {
	fmt.Println("wait groups in golang")
	websiteList := []string {
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	//below for loop iterates through the web list and store each item in web variable
	for _, web := range websiteList {
		//below two lines should be in the same order as written
		//if the order will be reverse, then we will not get any output except error
		//this is because waiting time will be 0 for the execution of go routine 
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()	//it will wait for the complete execution of the go routine
}

func getStatusCode(endpoint string) {
	defer wg.Done()	//this method will tell if the go routine has been successfully executed 

	res, err := http.Get(endpoint)

	if err != nil{
		fmt.Println("OOPS in status code")
	} else{

		//mutex is mutually exclusive lock. it locks a thread to perform its operations and after successful operation, this thread is unlocked by the mutex to give chance to other thread.
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}