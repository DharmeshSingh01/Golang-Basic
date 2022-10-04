package main

import "fmt"

func main() {
	defer fmt.Println("This call from defer1")
	defer fmt.Println("This call from defer2")
	defer fmt.Println("This call from defer3")
	defer fmt.Println("This call from defer4")
	fmt.Println("Defer in Golang")
	countLoop()

}

func countLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
