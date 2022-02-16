package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Intro")

	var fruits = []string{"Apple", "Banana"}

	fruits = append(fruits, "Mango", "Orange")

	fmt.Println("Slice elements are: ", fruits)
	fmt.Println("Length of slice is: ", len(fruits))

	fruits = append(fruits[1:])

	fmt.Println("Slice elements are: ", fruits)

	var nums = []int{6,8,1,2,0,3}
	sort.Ints(nums)
	fmt.Println("Slice elements are: ", nums)


	// Remove one element using slice
	index := 2

	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println(nums)
}