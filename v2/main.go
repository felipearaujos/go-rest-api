package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

var dao = ArticleDAO{}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	article.Id = bson.NewObjectId()
	err = dao.Insert(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := dao.ListAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(articles)
}

func init() {
	dao.Connect()
}

func main() {
	handlerRequest()
}
