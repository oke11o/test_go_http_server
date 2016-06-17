package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleIndex)
	fmt.Println(http.ListenAndServe(":10888", nil))
}

// HandleIndex is main route
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world")
}
