package main

import "fmt"

type User struct {
	FirstName string
	Age       int
}

//Go methods are similar to Go function with one difference,
//i.e, the method contains a receiver argument in it. With the help of the receiver argument, the method can access the properties of the receiver.
//Here, the receiver can be of struct type or non-struct type.

// func receiver name, receiver type, fnname, return type

// normal receiver
func (u1 User) printUserName() (string, int) {
	return (u1.FirstName), (u1.Age)
}

// pointer receiver
// With the help of a pointer receiver, if a change is made in the method, it will reflect in the caller which is not possible with the value receiver methods.
func (u1 *User) changeUserDetails(newname string, newage int) {
	(*u1).FirstName = newname
	(*u1).Age = newage
}

func main() {
	user1 := User{
		FirstName: "sudha",
		Age:       26,
	}
	fmt.Println(user1.printUserName())

	//pointer receiver
	p := &user1
	p.changeUserDetails("geetha", 45)
	fmt.Println(user1.printUserName())
}
