package product

import (
	"github.com/globalsign/mgo/bson"
)

type Product struct {
	ID      bson.ObjectId `json:"id"`
	Name    string        `json:"name"`
	Expired string        `json:"expired"`
}
