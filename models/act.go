package models

//import "gopkg.in/mgo.v2/bson"

type act struct {
	Account       string          `bson:"account" json:"account"`
	jName         string          `bson:"name"	  json:"name"`
	Authorization []authorization `bson:"authorization" json:"authorization"`
	Data          data            `bson:"data" json:"data"`
}
type data struct {
	From     string `bson:"from" json:"from"`
	To       string `bson:"to" json:"to"`
	Quantity string `bson:"quantity" json:"quantity"`
	Memo     string `bson:"memo" json:"memo"`
}
