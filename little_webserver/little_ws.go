package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Samwell!\n")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
