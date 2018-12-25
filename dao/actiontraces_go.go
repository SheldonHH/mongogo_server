package dao

import (
	"log"

	. "github.com/sheldonhh/mongogo_server/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ActiontracesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "action_traces"
)

// Establish a connection to database
func (m *ActiontracesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}


// Find list of movies
func (m *ActiontracesDAO) FindAll() ([]Act, error) {
	var acts []Act
	err := db.C(COLLECTION).Find(bson.M{}).All(&acts)
	return acts, err
}

// Find a act by its id
func (m *ActiontracesDAO) FindById(id string) (Act, error) {
	var act Act
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&act)
	return act, err
}

// Insert a act into database
func (m *ActiontracesDAO) Insert(act Act) error {
	err := db.C(COLLECTION).Insert(&act)
	return err
}

// Delete an existing act
func (m *ActiontracesDAO) Delete(act Act) error {
	err := db.C(COLLECTION).Remove(&act)
	return err
}

// Update an existing act
func (m *ActiontracesDAO) Update(act Act) error {
	err := db.C(COLLECTION).UpdateId(act.ID, &act)
	return err
}
