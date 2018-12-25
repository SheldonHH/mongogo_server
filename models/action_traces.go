package models
import "gopkg.in/mgo.v2/bson"

type action_traces struct {
	ID          	bson.ObjectId  `bson:"_id" json:"id"`
	Recepit        	recepit        `bson:"recepit" json:"recepit"`
	Act       		act			   `bson:"act" json:"act"`
	ContextFree 	bool		   `bson:"context_free" json:"context_free"`
	elasped 	    int            `bson:"elasped" json:"elasped"`
	console 		string     	   `bson:"console" json:"console"`
	TrxId 			string         `bson:"trx_id" json:"trx_id"`
	BlockNum 		int       	   `bson:"block_num" json:"block_num"`
	BlockTime		string 		   `bson:"block_time" json:"block_time`
	ProducerBlockId string		   `bson:"producer_block_id" json:"producer_block_id`
	AccountRamDeltas string		   `bson:"account_ram_deltas" json:"account_ram_deltas`
	Except 			string		   `bson:"except" json:"except`
	TrxStatus 		string	       `bson:"trx_status" json:"trx_status`
	CreatedAt 		string	       `bson:"created_at" json:"created_at`
}

