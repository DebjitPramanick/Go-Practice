package main

import "fmt"

const GlobalToken string = "asd34hj" // Public

func main(){
	var name string = "Debjit"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name) // Prints type

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallFloat float32 = 156.45
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var defInt int // Default value will be assigned
	fmt.Println(defInt)
	fmt.Printf("Variable is of type: %T \n", defInt)

	var impVar = "Random string" // Implicit type
	fmt.Println(impVar)
	fmt.Printf("Variable is of type: %T \n", impVar)

	noVarStyle := 30.5 // No Var Style (We cannot do this outside method/globally)
	fmt.Println(noVarStyle)
	fmt.Printf("Variable is of type: %T \n", noVarStyle)

	fmt.Println(GlobalToken)
	fmt.Printf("Variable is of type: %T \n", GlobalToken)
}