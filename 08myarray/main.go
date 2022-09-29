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
	myArray := [3]float64{1.2, 5.6}
	fmt.Println(myArray)
	myArray[2] = 6

	//a := 10
	// myArray[0] = a //cannot use a (variable of type int) as float64 value in

	//myArray[3] = 10.10 Index out of bound

	fmt.Println(myArray)

}
