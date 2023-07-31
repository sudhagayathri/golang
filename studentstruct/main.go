package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name      string
	Studentid int32
	Class     int32
}

func main() {
	std1 := Student{
		Name:      "sudha",
		Studentid: 1,
		Class:     3,
	}
	std2 := Student{
		Name:      "gayatri",
		Studentid: 2,
		Class:     3,
	}

	var response []Student
	response = append(response, std1)
	response = append(response, std2)
	data, _ := json.Marshal(response)
	fmt.Println(string(data))
}
