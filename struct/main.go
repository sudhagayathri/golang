package main

import (
	"log"
	"time"
)

// struct - user defined type which allows to combine items of different type to single type
// it is generally classes in object oritented approach
// it doesnt suport inheritance but supports composition

//suppose if we have user properties and want to define under user

type User struct {
	FirstName string
	LastName  string
	Age       int
	Birthdate time.Time
}

func main() {
	u1 := User{
		FirstName: "sudha",
		LastName:  "gayatri",
	}
	log.Println("name is ", u1.FirstName, u1.LastName, u1.Birthdate)
}
