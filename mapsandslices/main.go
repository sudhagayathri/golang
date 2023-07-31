package main

import (
	"log"
	"sort"
)

// map is kind of hashmaps - key, value
func main() {
	mymap := make(map[string]int)
	mymap["age"] = 15
	mymap["height"] = 6
	log.Println(mymap)

	mymap2 := map[int]string{
		41: "home",
		42: "details",
	}

	log.Println(mymap2[41])

	type User3 struct {
		FirstName string
		LastName  string
	}

	mymap3 := make(map[string]User3)

	me := User3{
		FirstName: "sudha",
		LastName:  "gayathri",
	}

	mymap3["res"] = me
	log.Println(mymap3["res"].FirstName)

	//slices - Slices in Go are a flexible and efficient way to represent arrays,
	//and they are often used in place of arrays because of their dynamic size and added features. A slice is a reference to a portion of an array. It’s a data structure that describes a portion of an array by specifying the starting index and the length of the portion

	var slicearr []string
	slicearr = append(slicearr, "experts", "trains", "hotels")
	log.Println(slicearr)

	slicearr2 := []int{4, 5, 1, 2, 3}
	slicearr2 = append(slicearr2, 9, 10)
	sort.Ints(slicearr2)
	log.Println(slicearr2)
	log.Println(slicearr2[0:4])

	//looping over arrays
	for _, verticals := range slicearr {
		log.Println(verticals)
	}
	//looping over map datastructure
	newobj := make(map[string]string)

	newobj["team"] = "xpertsnew"
	newobj["teamsize"] = "4"

	for datatype, data := range newobj {
		log.Println(datatype, data)
	}

	//creating array of structs
	type Userstruct struct {
		Name     string
		Age      int
		Category string
	}
	var users []Userstruct
	users = append(users, Userstruct{"Sudha", 26, "xperts"}, Userstruct{"madhu", 25, "xperts"})

	for _, data := range users {
		log.Println(data.Name)
	}
}
