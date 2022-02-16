package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps Intro")

	languages := make(map[string] int)

	languages["JS"] = 1
	languages["Python"] = 1
	languages["React"] = 2
	languages["Node"] = 1

	fmt.Println(languages)

	fmt.Println(languages["JS"])

	delete(languages, "JS")
	fmt.Println(languages)

	// Loop through map

	for key, val:=range languages{
		fmt.Printf("%v - %v\n", key, val)
	}
}