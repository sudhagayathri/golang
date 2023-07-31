package main

import (
	"fmt"
	"log"

	"github.com/sudhagayathri/golang/helpers"
)

func main() {
	//l1- variables and functions
	fmt.Println("hello go")
	var txtstr string
	txtstr = "hi mobileapi"
	fmt.Println(txtstr)

	var i int
	i = 7
	fmt.Println("i is assigned to", i)
	var whatsay string
	whatsay = saySomething()
	fmt.Println("this is said ", whatsay)
	//other way to call fn
	whatsay2 := saySomething()
	fmt.Println("this is new said ", whatsay2)

	whatsay3, whatsay4 := sayMore()
	fmt.Println("this is other new said ", whatsay3, whatsay4)

	//var helpervar helpers.CompanyDetails
	helpervar := helpers.CompanyDetails{CompanyName: "xcc"}
	log.Println(helpervar.CompanyName)
	log.Println(helpers.PrintCompanyDetails())
}

func saySomething() string {
	return "said hello"
}

// returning more than one param
func sayMore() (string, string) {
	return "something", "else"
}
