package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// using range
	for i := range days {
		fmt.Println(days[i])
	}

	// like for each
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}
	// if want to ignore index value
	for _, day1 := range days {
		fmt.Printf("Index is  and value is %v\n", day1)
	}
	rougueValue := 1
	for rougueValue <= 10 {
		if rougueValue == 2 {
			goto ds
		}
		if rougueValue == 4 {
			rougueValue++
			fmt.Println("Continue is workin")
			continue

		}
		if rougueValue == 6 {
			fmt.Println("Break stement Work")
			break
		}
		fmt.Println(rougueValue)
		rougueValue++
	}
ds:
	fmt.Println("jumping at lable")
}
