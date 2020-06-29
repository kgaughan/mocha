package main

import (
	"io"
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", Root)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
