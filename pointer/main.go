package main

import "log"

func main() {
	//pointer - variable storing reference of another variable
	var p string = "sudha"
	//pointer declaration
	var x *string //here it is pointer for string variable
	x = &p
	log.Println("value of p ", p)
	log.Println("value of &p ", &p)
	log.Println("value of x ", x)
	log.Println("value stored at x -derference of pointer ", *x)
	newvalue := "gayathri"
	//changing content of that location
	*x = newvalue
	log.Println("new val of p ", p)
	//pointer using fn
	s := "newsudha"
	log.Println("value of s is ", s)
	changeValueusingPointer(&s)
	log.Println("value of s after pointer is ", s)
}

func changeValueusingPointer(s *string) {
	newstr := "newgayatri"
	*s = newstr
}
