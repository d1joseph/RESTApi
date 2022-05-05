// A simple server to handle HTTP requests
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article struct
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// global Article array
var Articles []Article

// homePage returns the index of the server
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page.")
	fmt.Println("Endpoint Hit: homePage")
}

// handleRequests listens at specified port and maps
// the URL path hit with the approriate request handler
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// index
	myRouter.HandleFunc("/", homePage)

	// /articles
	myRouter.HandleFunc("/articles", returnAllArticles)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// returnAllArticles returns all article data available
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article description", Content: "Article content"},
		Article{Title: "Hello again", Desc: "Another article description", Content: "Latest article content"},
	}

	handleRequests()
}
