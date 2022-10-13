package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://courses.learncodeonline.in/learn/Python-Mock-Test"

func main() {
	fmt.Println("Handling url in Golang")
	fmt.Println(myurl)
	//parsing
	fmt.Println("tests")
	result, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("test")
		panic(err)
	}

	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)

	//createurl

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)
}
