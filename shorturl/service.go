package shorturl

import (
	"tiny_url/utils"
	"sync"
	"strings"
)

var longUrlMap = new(sync.Map)
var shortUrlMap = new(sync.Map)

type service struct {
	Repo Repository
	RedisClient utils.RedisClient
}

type Service interface {
	Create(entry UrlEntry) (CreateResponse,error)
	Get(shortUrl string) (CreateResponse,error)
}

type ErrorMsg struct {
	Message string
}

type CreateResponse struct {
	ShortUrl string
	LongUrl string
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
	return CreateResponse{LongUrl:created.LongUrl, ShortUrl:created.ShortUrl},nil
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
	return CreateResponse{LongUrl:entry.LongUrl, ShortUrl:entry.ShortUrl},nil
}


//NewService is abstract service implementation for db calls
func NewService(repo Repository, redisClient utils.RedisClient) Service {
	return &service{Repo: repo, RedisClient:redisClient}
}

