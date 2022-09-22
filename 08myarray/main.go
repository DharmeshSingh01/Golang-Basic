package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golangs")
	// Creating  an Array
	var fruiteList [4]string
	// Intialise Array
	fruiteList[0] = "Apple"
	fruiteList[1] = "Orange"
	fruiteList[3] = "Mango"

	fmt.Println("Fruit List is :", fruiteList)
	fmt.Println("Lenth of Fruit List is: ", len(fruiteList))

	// create array and intialise value at same time

	var vegList = [3]string{"Potato", "Beens", "Mashrom"}
	fmt.Println("Veg List is :", vegList)
	fmt.Println("Length of vegList List is: ", len(vegList))

}
