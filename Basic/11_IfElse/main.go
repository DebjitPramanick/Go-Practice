package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	count := 10

	// General syntax
	if count < 10 {
		fmt.Println("Allowed")
	} else {
		fmt.Println("Not allowed")
	}

	// Special syntax

	if num := 3; num < 10 {
		fmt.Println("Allowed")
	} else {
		fmt.Println("Not allowed")
	}
}
