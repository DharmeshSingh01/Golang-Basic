package main

import "fmt"

func main() {

	// creating pointer
	var ptr *int

	fmt.Println("pointer created and it value :", ptr)
	// assigning pointer
	mynumber := 23
	var ptr2 = &mynumber
	fmt.Println("Pointer value is", ptr2)
	fmt.Println("Value point by pointer is ", *ptr2)

	// change original value of variable using pointer

	*ptr2 = *ptr2 + 2
	fmt.Println("Value change by pointer is ", *ptr2)

}
