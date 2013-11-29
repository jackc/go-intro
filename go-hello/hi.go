package main

import (
	"io"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world")
}

func main() {
	http.HandleFunc("/hi", hiHandler)
	http.ListenAndServe(":8080", nil)
}
