package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Caption     string        `json:"caption" bson:"caption"`
	Image       string        `json:"image" bson:"image"`
	CreatedDate time.Time     `json:"createdDaate"`
	Owner       string        `json:"owner" bson:"owner"`
}
