package main

import "fmt"

func main() {
	fmt.Println("Pointers Intro")
	var ptr *int

	var val int = 9
	ptr = &val

	fmt.Println("Address stored in pointer is: ", ptr)
	fmt.Println("Value stored in pointer is: ", *ptr)
}