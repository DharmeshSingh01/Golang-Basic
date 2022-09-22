package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Number Of CPU in you System:", runtime.NumCPU())
	fmt.Println("OS :", runtime.GOOS)

	fmt.Println("Current Time is:")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.Format("02-01-2006"))

	// create custome date

	createdDate := time.Date(2022, time.January, 20, 23, 23, 0, 0, time.UTC)
	fmt.Println("Created Date is :- ", createdDate.Format("02-01-2006 Monday"))

}
