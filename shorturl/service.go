package shorturl

import (
	"tiny_url/utils"
	"sync"
	"strings"
)

var longUrlMap = new(sync.Map)
var shortUrlMap = new(sync.Map)

const (
	//Uable to retrieve the request with host url when running locally.
	//It works when running on any external server. but for now hardcoding this. 
	//Need to refactor this
	tinyurlhosturi = "http://localhost:8000/tinyurl/"
)
type service struct {
	Repo Repository
	RedisClient utils.RedisClient
}

type Service interface {
	Create(entry UrlEntry) (CreateResponse,error)
	Get(shortUrl string) (CreateResponse,error)
}

type ErrorMsg struct {
	Message string `json:"message,omitempty"`
}

type CreateResponse struct {
	ShortUrl string `json:"shortUrl"`
	LongUrl string `json:"longUrl"`
	ErrorMsg
}

func (s *service) Create(entry UrlEntry) (CreateResponse,error) {
	var created UrlEntry
	var err error

	tObj, ok := longUrlMap.Load(entry.LongUrl)
	if !ok {
		created,err = s.Repo.FindByLongurl(entry.LongUrl)
		if(created.LongUrl == "" || (err != nil && strings.Trim(err.Error(), "") == "not found")) {
			created, err = s.Repo.Create(entry)
			if (err != nil) {
				return CreateResponse{}, err
			}
		}
		longUrlMap.Store(entry.LongUrl, created)
		shortUrlMap.Store(created.ShortUrl,created)
	} else {
		created = tObj.(UrlEntry)
	}
	return CreateResponse{LongUrl:created.LongUrl, ShortUrl:tinyurlhosturi + created.ShortUrl},nil
}

func (s *service) Get(shortUrl string) (CreateResponse,error) {
	var entry UrlEntry
	var err error

	tObj, ok := shortUrlMap.Load(shortUrl)
	if !ok {
		entry,err = s.Repo.Get(shortUrl)
		if(err != nil || entry.ShortUrl == "") {
			return CreateResponse{ErrorMsg:ErrorMsg{Message:utils.ErrNotFound.Error()}},utils.ErrNotFound
		}
		shortUrlMap.Store(entry.ShortUrl, entry)
	} else {
		entry = tObj.(UrlEntry)
	}
	return CreateResponse{LongUrl:entry.LongUrl, ShortUrl:tinyurlhosturi+entry.ShortUrl},nil
}


//NewService is abstract service implementation for db calls
func NewService(repo Repository, redisClient utils.RedisClient) Service {
	return &service{Repo: repo, RedisClient:redisClient}
}

