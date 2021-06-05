package log

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type RequestInfo struct {
	Length  int    `bson:"length" json:"length"`
	Method  string `bson:"method" json:"method"`
	URI     string `bson:"uri" json:"uri"`
	Headers bson.M `bson:"headers" json:"-"`
}

type ResponseInfo struct {
	Length  int    `bson:"length" json:"length"`
	Status  int    `bson:"status" json:"status"`
	Headers bson.M `bson:"headers" json:"-"`
}

type Record struct {
	Id        uuid.UUID     `bson:"_id" json:"-"`
	Client    string        `bson:"client" json:"client"`
	Remote    string        `bson:"remote_addr" json:"remote_addr"`
	Service   string        `bson:"service" json:"service"`
	Unknown   bool          `bson:"unknown" json:"unknown"`
	Forbidden bool          `bson:"forbidden" json:"forbidden"`
	Request   *RequestInfo  `bson:"request" json:"request"`
	Response  *ResponseInfo `bson:"response" json:"response"`
}
