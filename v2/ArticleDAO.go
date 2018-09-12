package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type ArticleDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (COLLECTION = "Articles")

func (m *ArticleDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ArticleDAO) Insert(article Article) error {
	err := db.C(COLLECTION).Insert(&article)
	return err
}

