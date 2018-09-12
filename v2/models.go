package main

import "gopkg.in/mgo.v2/bson"

type Article struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Title   string        `bson:"Title" json:"Title"`
	Desc    string        `bson:"Desc" json:"Desc"`
	Content string        `bson:"Content" json:"Content"`
}

type Articles []Article
