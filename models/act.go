package models
//import "gopkg.in/mgo.v2/bson"

type act struct {
	Account         string        `bson:"account" json:"account"`
	jName  			string        `bson:"name"	  json:"name"`
	Authorization   authorization `bson:"authorization" json:"authorization"`
	Data 			string        `bson:"data" json:"data"`
}