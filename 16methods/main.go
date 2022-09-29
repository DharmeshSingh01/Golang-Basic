package main

import "fmt"

func main() {
	fmt.Println("Method in Golang ")
	dharmesh := User{"Dharmesh", "ds29@go.dev", true, 0}

	dharmesh.GetStatus()
	dharmesh.GetNewMail()

	fmt.Printf("User Name is %v and Email is :- %v", dharmesh.Name, dharmesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int // this is not exportable
}

func (u User) GetStatus() {
	fmt.Println("Is User active : ", u.Status)
}
func (u User) GetNewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New Email Id is :- ", u.Email)
}
