package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File in Golang")
	content := "This beed to go in a file"
	file, err := os.Create("./mylogfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkError()
	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkError(err)
	fmt.Println("Length is", length)
	defer file.Close()
	readFile("./mylogfile.txt")
}
func readFile(filename string) {
	dataBytes, err := os.ReadFile(filename)
	checkError(err)
	fmt.Println("File Data in Bytes", dataBytes)
	fmt.Println("Text inside a file:- ", string(dataBytes))
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
