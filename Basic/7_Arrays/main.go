package main

import "fmt"

func main() {
	fmt.Println("Array Intro")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Tomato"

	fmt.Println("Array elements are: ", fruits)
	fmt.Println("Length of array is: ", len(fruits))

	var vegs = [3]string{"Carrot", "Potato","Ginger"}

	fmt.Println("Array elements are: ", vegs)
	fmt.Println("Length of array is: ", len(vegs))
}