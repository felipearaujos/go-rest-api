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
	myRouter.HandleFunc("/articles", createArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
