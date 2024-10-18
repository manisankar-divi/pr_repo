package main

import (
	"fmt"
	"net/http"
)

// helloHandler handles HTTP requests to the root URL.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// This will return the string without a newline
	fmt.Fprint(w, "Hello, World!") // Change Fprintln to Fprint
}

// main function sets up the HTTP server.
func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
