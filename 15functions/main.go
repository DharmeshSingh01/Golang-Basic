package main

import "fmt"

func main() {
	fmt.Println("Function in Golang")
	greeter()

	/*
		// we can't define a function inside a function
		func secondFunction()  {
			fmt.Println("Second method ")
		}
		secondFunction()
	*/
	valadd := adder(2, 4)
	fmt.Println(valadd)

	proResult := proAdder(2, 23, 56, 12)
	fmt.Println("proresult is", proResult)

	fmt.Println("Return more than one value from function")

	multival, msg := multi(3, 6)
	fmt.Println(msg, multival)

}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

// ..verdic
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// return more than one value from function

func multi(val1 int, val2 int) (int, string) {

	multiple := val1 * val2
	return multiple, "value are multiplied"

}

func greeter() {
	fmt.Println("Hello from Function ")
}
