package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")
	loginUser := 23
	var loginResult string
	if loginUser < 10 {
		loginResult = "Regular User"
	} else if loginUser > 10 {
		loginResult = "Watch Out"
	} else {
		loginResult = "Exactly 10 User"
	}
	fmt.Println(loginResult)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}
	if num := 3; num > 10 {
		fmt.Println("Num is Greter than 10")
	} else {
		fmt.Println("Num is less than 10")
	}

}
