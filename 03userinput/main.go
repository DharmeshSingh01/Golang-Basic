package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enyter the rating for our Pizza")

	// comma ok OR comma error

	input, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)
	fmt.Println("Thanks for rating", err)

}
