package main

import "fmt"

func main() {

	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@gamil.com", 01734, false, 35}

	fmt.Println(hitesh)
	fmt.Printf("Detail of %+v\n", hitesh)
	fmt.Printf("Name is: %v and email is: %v\n", hitesh.Name, hitesh.Email)

	hitesh.GetStatus()
	hitesh.NewEmain()

}

type User struct {
	Name   string
	Email  string
	Number int
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user actice: ", u.Status)
}

func (u User) NewEmain() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user: ", u.Email)
}
