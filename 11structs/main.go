package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	dharmesh := User{"DHarmesh", "dharmeshsingh@live.com", true, 30}
	fmt.Println(dharmesh)

	fmt.Printf("Users deatails are : %+v\n", dharmesh)

	fmt.Printf("Name is %v and email is %v", dharmesh.Name, dharmesh.Email)
}

// Define Struct

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
