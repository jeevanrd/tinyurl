package shorturl

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION = "url"
)

type repository struct {
	Db *mgo.Database
}

type Repository interface {
	Get(tinyUrl string) (UrlEntry,error)
	Create(entry UrlEntry) (UrlEntry, error)
	FindByLongurl(longurl string) (UrlEntry,error)
}

func (r *repository) Create(entry UrlEntry) (UrlEntry, error) {
	entry.ID = bson.NewObjectId()
	entry.ShortUrl = entry.ID.Hex()
	err := r.Db.C(COLLECTION).Insert(&entry)
	return entry,err
}

func (r *repository) Get(tinyUrl string) (UrlEntry,error) {
	var url UrlEntry
	err := r.Db.C(COLLECTION).Find(bson.M{"shortUrl":tinyUrl}).One(&url)
	return url,err
}

func (r *repository) FindByLongurl(longurl string) (UrlEntry,error) {
	var url UrlEntry
	err := r.Db.C(COLLECTION).Find(bson.M{"longUrl":longurl}).One(&url)
	return url,err
}

func NewRepository(Db *mgo.Database) Repository {
	return &repository{Db: Db}
}
