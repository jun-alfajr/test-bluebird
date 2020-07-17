package product

import (
	"github.com/globalsign/mgo/bson"
)

type Product struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Expired string        `json:"expired" bson:"expired"`
}
