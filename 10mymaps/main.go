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
	fmt.Println("List of All Language", language)
	fmt.Println("Js shorts for", language["JS"])
}
