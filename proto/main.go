package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	sudha := &Personproto{
		Name: "Sudha",
		Age:  26,
	}

	data, err := proto.Marshal(sudha)
	fmt.Println(err)

	fmt.Println(data)
}
