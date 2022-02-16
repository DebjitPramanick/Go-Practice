package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in Golang")
	debjit := User{"Debjit", "deb@gmail.com", true, 21}

	fmt.Println(debjit)

	fmt.Printf("Details: %+v\n", debjit)

	fmt.Printf("Name is: %v", debjit.Name)
}
