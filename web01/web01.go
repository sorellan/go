package main

import (
	"fmt"
	"net/http"
)

var closed = true

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Samwell!</h1><p>¿La cerradura esta cerrada? %t</p>", closed)
}

func lock(w http.ResponseWriter, req *http.Request) {
	if closed {
		fmt.Fprintf(w, "Ya estaba cerrada\n")
	} else {
		closed = true
		fmt.Fprintf(w, "Cerrada satisfactoriamente\n")
	}
}

func unlock(w http.ResponseWriter, req *http.Request) {
	if closed {
		closed = false
		fmt.Fprintf(w, "Abierta satisfactoriamente")
	} else {
		fmt.Fprintf(w, "Ya estaba abierta")
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/close", lock)
	http.HandleFunc("/open", unlock)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
