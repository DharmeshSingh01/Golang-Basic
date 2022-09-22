package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Input Conversion")
	fmt.Println("Enter your rating for our PIZZA!")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your Rating", input)

	fmt.Println("we will add 1 to your rating")

	numinput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	//numinput, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your Rating after ADD 1:", numinput+1)
	}

}
