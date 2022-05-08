// A simple server to handle HTTP requests
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article struct
type Article struct {
	Id      string `jason:"Id"`
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

// handleRequests listens at the specified port and maps
// the URL path hit with the approriate request handler to call
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// index
	myRouter.HandleFunc("/", homePage)
	// /articles
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	// listens on localhost port int
	portNumber := fmt.Sprintf(":%v", 3000)
	fmt.Println("Listening on port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, myRouter))
}

// returnAllArticles returns all article data available
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// returnSingleArticle returns a single article
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// createNewArticle
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of the POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article description", Content: "Article content"},
		Article{Id: "2", Title: "Hello again", Desc: "Another article description", Content: "Latest article content"},
	}
	handleRequests()
}
