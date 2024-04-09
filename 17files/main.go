package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "Content to write in golang file"

	file, err := os.Create("./myGolangFile.txt")

	// if err!=nil{
	// 	panic(err)
	// }

	checkErr(err)

	length, err := io.WriteString(file, content)
	fmt.Println("The length of file is:", length)

	defer file.Close()
	readFile("./myGolangFile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }

	checkErr(err)

	fmt.Println("The content of the file is: \n", databyte)
	fmt.Println("------------")
	fmt.Println("The content of the file is: \n", string(databyte))
}


func checkErr(err error){
	if err != nil {
		panic(err)
	}
}