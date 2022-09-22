package main

import "fmt"

const LoginToken = "dharmesh123"

func main() {
	//fmt.Println("Variables")

	var username string = "Dharmesh"
	fmt.Println(username)
	fmt.Printf("Variable is of Type : %T \n", username)

	var boolvalue bool = true
	fmt.Println(boolvalue)
	fmt.Printf("Variable is of Type : %T \n", boolvalue)

	// 0 to 255
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of Type : %T \n", smallValue)

	//implicit Type

	var website = "learncodeonline.com"
	fmt.Println(website)
	// no var style
	numberOfUser := 300
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable type is: %T \n", LoginToken)

}
