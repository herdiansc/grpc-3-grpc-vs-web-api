package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("incoming\n")
		fmt.Fprintf(w, "hello world")
	})

	http.ListenAndServe(":8090", nil)
}
