package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ArticleId   string `json:Title`
	Description string `json:description`
	Content     string `json:content`
}

var articles map[string]Article

func main() {
	initialize()
	handleRequests()
}

func initialize() {
	articles = map[string]Article{
		"Go Lang Routines":    Article{"Go Lang Routines", "About go routines", "blah blah blah"},
		"Java MultiThreading": Article{"Java MultiThreading", "concurrency and synchronization in java", "blah ... blah"}}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", respondToRequests)
	router.HandleFunc("/Article/All", getAllArticles)
	router.HandleFunc("/Article/{id}", getArticleById).Methods("GET")
	router.HandleFunc("/Article", createArticle).Methods("POST")
	router.HandleFunc("/Article/All", deleteArticle).Methods("DELETE")
	router.HandleFunc("/Article/{id}", deleteArticleById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func respondToRequests(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Application is up and Running")
}

func getAllArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Request received for sending all articles")

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	var articleArray []Article = make([]Article, len(articles))
	var index = 0
	for _, value := range articles {
		articleArray[index] = value
		index++
	}

	art, _ := json.Marshal(articleArray)
	writer.Write(art)
}

func getArticleById(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(req)
	id := vars["id"]

	if art, found := articles[id]; found {
		w.WriteHeader(http.StatusAccepted)
		jsonValue, _ := json.Marshal(art)
		w.Write(jsonValue)
	} else {
		w.Write([]byte("Article not found"))
		w.WriteHeader(http.StatusNotFound)
	}
}

func createArticle(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	jsonValue, _ := json.Marshal(reqBody)

	fmt.Println(jsonValue)
}

func deleteArticle(w http.ResponseWriter, req *http.Request) {
	for k := range articles {
		delete(articles, k)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All articles deleted"))
}

func deleteArticleById(w http.ResponseWriter, req *http.Request) {

	key := mux.Vars(req)["id"]

	delete(articles, key)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Article deleted"))
}
