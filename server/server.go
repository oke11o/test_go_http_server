package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleIndex)
	fmt.Println(http.ListenAndServe(":10888", nil))
}

var gCounter int

// HandleIndex is main route
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	gCounter++
	// fmt.Printf("\r%d", gCounter)
	fmt.Fprintln(w, "hello, world")
}
