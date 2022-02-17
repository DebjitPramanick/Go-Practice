package main
import "fmt"

func showName(name string) {
	fmt.Printf("Hello %v!\n", name)
}

func add(a int, b int) int {
	return a+b
}

func proAdder(vals ...int) int {
	total := 0

	for _,val:=range vals {total += val}
	return total
}

func mulTypeReturn(vals ...int) (int, string) {
	total := 1

	for _,val:=range vals {total *= val}
	return total, "Result generated successfully."
}

func main(){
	fmt.Println("Hello World")
	showName("Debjit")
	sum := add(4,5)
	fmt.Printf("The sum is: %v\n", sum)

	proRes := proAdder(3,4,5,7,3)
	fmt.Printf("Sum of values is: %v\n", proRes)

	mulRes, message := mulTypeReturn(3,5,1,2,4)
	fmt.Println(message)
	fmt.Printf("Multiplication of values is: %v\n", mulRes)
}