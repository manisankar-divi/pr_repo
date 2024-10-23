package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// init initializes the logger
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

// helloHandler handles HTTP requests to the root URL.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

// startServer starts the HTTP server.
func startServer() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

// main function
func main() {
	startServer()
}
