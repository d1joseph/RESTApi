// A simple server to handle HTTP requests
package main

import (
	"fmt"
	"log"
	"net/http"
)

// homePage returns the index of the server
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page.")
	fmt.Println("Endpoint Hit: homePage")
}

// handleRequests listens at specified port and serves
// the URL path hit with the request handler
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
