package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")
	language := make(map[string]string)
	language["JS"] = "javascript"
	language["RE"] = "Reactjs"
	language["PY"] = "Python"
	language["RB"] = "Ruby"
	fmt.Println("List of All Language", language)
	fmt.Println("Js shorts for", language["JS"])

	fmt.Println("Deleteing value from map ")
	// Deleteing
	delete(language, "RB")
	fmt.Println("List of All Language after delete", language)

	// loops are intresing in Golang

	for key, value := range language {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// other syntax is comma err,
	for _, val := range language {
		fmt.Printf("Value is %v\n", val)
	}
}
