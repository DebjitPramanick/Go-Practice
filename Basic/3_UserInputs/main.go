package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome:="Hello User!"
	fmt.Println(welcome)

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok || error ok syntax

	input, _ := reader.ReadString('\n') // Input will take the input and _ will take any error
	fmt.Println("Your name is: ", input)
}