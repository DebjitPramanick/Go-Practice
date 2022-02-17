package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thur", "Fri", "Sat"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Printf("Day %v: %v\n", d+1, days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i, day := range days {
		fmt.Println(i, day)
	}

	val := 1

	for val < 10 {
		if val == 2 {
			goto lco
		}
		if val == 5 {
			val++
			continue
		}
		fmt.Println("Valus is: ", val)
		val++
	}

lco:
	fmt.Println("Jumping at this label") // Label
}
