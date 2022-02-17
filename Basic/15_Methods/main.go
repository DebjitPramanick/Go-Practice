package main
import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){ // If this method will be exported, make first letter capital
	fmt.Printf("Status is: %v", u.Status)
}

func (u* User) ChangeEmail(email string){
	u.Email = email
	fmt.Printf("Updated email is: %v\n", u.Email)
}

func main(){
	fmt.Println("Hello World")
	debjit := User{"Debjit", "deb@gmail.com", true, 21}
	debjit.GetStatus()
	debjit.ChangeEmail("deb1@gmail.com")
	fmt.Println(debjit)
}