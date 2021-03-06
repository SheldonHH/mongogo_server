package dao

import (
	"encoding/json"
	"fmt"
	. "github.com/sheldonhh/mongogo_server/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
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

// https://godoc.org/github.com/globalsign/mgo/bson#M
func (m *ActiontracesDAO) FindByQuery(query string) (ActionTraceResponse, error) {
	var act ActionTraceResponse

	//json.Marshal()
	// 	expect := `{"act.name":"transfer","act.account":"miaomiaobibi",$or:[{"act.data.to":"youbeforeme1"},{"act.data.from":"youbeforeme1"}]}`
	type FindReq struct {
		ActName    string              `json:"act.name"`
		ActAccount string              `json:"act.account"`
		Or         []map[string]string `json:"$or"`
	}
	var findReq = FindReq{
		ActAccount: "miaomiaobibi",
		ActName:    "transfer",
		Or: []map[string]string{
			{"act.data.to": "youbeforeme1"},
			{"act.data.from": "youbeforeme1"},
		},
	}

	b, err := json.Marshal(findReq)

	fmt.Printf("=======\n"+
		"b: %v, err: %v", string(b), err)

	//{"act.name":"transfer","act.account":"miaomiaobibi",$or:[{"act.data.to":"youbeforeme1"},{"act.data.from":"youbeforeme1"}]},{"act":1,"block_time":1}
	//db.action_traces.find({"act.name":"transfer","act.account":"miaomiaobibi",$or:[{"act.data.to":"youbeforeme1"},{"act.data.from":"youbeforeme1"}]})
	err = db.C(ACTCOLLECTION).Find(`{"act.name":"transfer","act.account":"miaomiaobibi",$or:[{"act.data.to":"youbeforeme1"},{"act.data.from":"youbeforeme1"}]}`).One(&act)
	err = db.C(ACTCOLLECTION).Find(
		bson.M{
			"act.name":    "transfer",
			"act.account": "miaomiaobibi",
			"$or":         []map[string]string{{"act.data.to": "youbeforeme1"}, {"act.data.from": "youbeforeme1"}},
		}).One(&act) //.Select(bson.M{"act", 1, "block_time": 1})
	if err != nil {
		fmt.Printf("err:%v", err)
	}

	//		bson.M{
	//			"reciever":"1",
	//			"account":"1",
	//			"name":"1",
	//		}).One(&act)
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
