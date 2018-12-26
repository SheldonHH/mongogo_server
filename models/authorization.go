package models

//import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type authorization struct {
	Actor      string `bson:"actor" json:"actor"`
	Permission string `bson:"permission" json:"permission"`
}
