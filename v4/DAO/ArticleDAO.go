package DAO

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2/bson"

	configs "../Config"
	models "../Models"
	mgo "gopkg.in/mgo.v2"
)

type ArticleDAO struct {
}

var db *mgo.Database

const COLLECTION = "golang.Articles"

func (m *ArticleDAO) Connect() {
	var config = configs.Config{}
	config.Read()

	session, err := mgo.Dial(config.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(config.Database)
}

func (m *ArticleDAO) Insert(article models.Article) error {
	err := db.C(COLLECTION).Insert(&article)
	return err
}

func (m *ArticleDAO) ListAll() (models.Articles, error) {
	var articles models.Articles
	err := db.C(COLLECTION).Find(bson.M{}).All(&articles)
	return articles, err
}

func (m *ArticleDAO) Get(id string) (models.Article, error) {
	var article models.Article

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
