package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	myRouter.HandleFunc("/articles", createArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
