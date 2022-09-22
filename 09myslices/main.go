package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	fruiteList := []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of FruitList is %T\n", fruiteList)
	fruiteList = append(fruiteList, "Mango", "Banana")
	fmt.Println(fruiteList)

	//Slice

	fruiteList = append(fruiteList[1:3])
	//fruiteList = append(fruiteList[:3]) //[Apple Tomato Peach]
	fmt.Println(fruiteList)

	highScores := make([]int, 4)
	highScores[0] = 856
	highScores[1] = 342
	highScores[2] = 500
	highScores[3] = 654
	//highScores[4]=777 // through an error Index out of Bound

	highScores = append(highScores, 555, 654, 999)
	fmt.Println(highScores)

	// sort
	fmt.Println("Is Array sorted ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println("Is Array sorted ", sort.IntsAreSorted(highScores))

	// How to remove a value from slices based on Index

	var courses = []string{"reactjs", "Javascript", "Swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After deleteing value from Index 2")
	fmt.Println(courses)
}
