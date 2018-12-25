package dao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ActiontracesDAO struct {
	Server   string
	Database string
}


const (
	ACTCOLLECTION = "action_traces"
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
func (m *ActiontracesDAO) FindAll() ([]ActionTrace, error) {
	var actiontraces []ActionTrace
	err := db.C(ACTCOLLECTION).Find(bson.M{}).All(&actiontraces)
	return actiontraces, err
}

// Find a act by its id
func (m *ActiontracesDAO) FindById(id string) (ActionTrace, error) {
	var act ActionTrace
	err := db.C(ACTCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&act)
	return act, err
}

// Insert a act into database
func (m *ActiontracesDAO) Insert(act ActionTrace) error {
	err := db.C(ACTCOLLECTION).Insert(&act)
	return err
}

// Delete an existing act
func (m *ActiontracesDAO) Delete(act ActionTrace) error {
	err := db.C(ACTCOLLECTION).Remove(&act)
	return err
}

// Update an existing act
func (m *ActiontracesDAO) Update(act ActionTrace) error {
	err := db.C(ACTCOLLECTION).UpdateId(act.ID, &act)
	return err
}
