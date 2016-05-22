package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkPassword(pwd string) bool {
	return pwd == "1234"
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("New connection from " + req.RemoteAddr + "\n")
	fmt.Printf("Method: " + req.Method + "\n")
	fmt.Printf("Password: " + req.FormValue("password") + "\n")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)

	checkPwd := checkPassword(req.FormValue("password"))
	if checkPwd {
		io.WriteString(w, "Valid password\n")
	} else {
		io.WriteString(w, "Invalid password\n")
	}
	//w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	//w.Header().Add("Trailer", "AtEnd3")

	//io.WriteString(w, "This HTTP response is a hello world.\n")
	//fmt.Fprintf(w, "This HTTP response has both headers before this text and trailers at the end.\n")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
