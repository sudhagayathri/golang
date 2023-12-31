package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloHandler)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Hello, there\n")
}
