package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("New connection from " + req.RemoteAddr + "\n")
	w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	w.Header().Add("Trailer", "AtEnd3")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, "This HTTP response is a hello world.\n")
	//fmt.Fprintf(w, "This HTTP response has both headers before this text and trailers at the end.\n")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
