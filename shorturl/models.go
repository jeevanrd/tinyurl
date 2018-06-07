package shorturl

import "gopkg.in/mgo.v2/bson"

type UrlEntry struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	LongUrl     string        `bson:"longUrl" json:"longUrl"`
	ShortUrl  	string        `bson:"shortUrl" json:"shortUrl"`
}