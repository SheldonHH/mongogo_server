package models
import "gopkg.in/mgo.v2/bson"

type receipt struct {
	Name        	string        `bson:"receiver" json:"receiver"`
	ActDigest       string		  `bson:"act_digest" json:"act_digest"`
	GlobalSequence 	int		      `bson:"global_sequence" json:"global_sequence"`
	RecvSequence 	string        `bson:"recv_sequence" json:"recv_sequence"`
	AuthSequence 	[][]string     `bson:"auth_sequence" json:"auth_sequence"`
	CodeSequence 	string        `bson:"code_sequence" json:"code_sequence"`
	AbiSequence 	string        `bson:"abi_sequence" json:"abi_sequence"`
}

