package storage

import (
	"gopkg.in/mgo.v2/bson"
)

type IP struct {
	ID       bson.ObjectId `bson:"_id" json:"-"`
	IP       string        `bson:"ip" json:"ip"`
	Port     int           `bson:"port" json:"port"`
	Protocol string        `bson:"protocol" json:"protocol"`
	Usable   bool          `bson:"usable" json:"usable"`
}

// NewIP .
func NewIP() *IP {
	return &IP{
		ID: bson.NewObjectId(),
	}
}
