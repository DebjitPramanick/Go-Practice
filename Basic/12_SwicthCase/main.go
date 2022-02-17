package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World")

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(6) + 1 // Range 0 - 5

	fmt.Println("Score: ", num)

	switch num { // use fallthrough to execute all other valid next cases
	case 1:
		fmt.Println("Poor Performance")
	case 2:
		fmt.Println("Not Satisfied")
	case 3:
		fmt.Println("Good")
	case 4:
		fmt.Println("Excellent")
	case 5:
		fmt.Println("Outstanding")
	default:
		fmt.Println("Invalid")
	}
}
