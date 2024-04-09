package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	punit := User{"Punit", "punitgoyal@gmail.com", true, 21}
	fmt.Println(punit)

	//both have same output
	fmt.Println("My details are:", punit)
	fmt.Printf("My details are: %v\n", punit)

	fmt.Printf("My details are: %+v\n", punit)

	//both have same output
	fmt.Printf("My name is %v and email is %v\n", punit.Name, punit.Email)
	fmt.Printf("My name is %+v and email is %+v\n", punit.Name, punit.Email)

	punit.GetStatus()
	punit.NewEmail()

	fmt.Println("The email of the user is:", punit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("The status of the user is:", u.Status)
}

func (u User) NewEmail(){
	u.Email = "punit_11212505@mmumullana.org"
	fmt.Println("The new email of the user is:", u.Email)
}