package main

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type ArticleDAO struct {
}

var db *mgo.Database

const COLLECTION = "golang.Articles"

func (m *ArticleDAO) Connect() {
	var config = Config{}
	config.Read()

	session, err := mgo.Dial(config.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(config.Database)
}

func (m *ArticleDAO) Insert(article Article) error {
	err := db.C(COLLECTION).Insert(&article)
	return err
}

func (m *ArticleDAO) ListAll() (Articles, error) {
	var articles Articles
	err := db.C(COLLECTION).Find(bson.M{}).All(&articles)
	return articles, err
}

func (m *ArticleDAO) Get(id string) (Article, error) {
	var article Article
	
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&article)
	return article, err
}

func (m *ArticleDAO) Remove(id string) error {

	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid Id")
	}

	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}
