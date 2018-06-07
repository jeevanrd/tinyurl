package shorturl

import (
	"tiny_url/utils"
	"fmt"
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
	message string
}

type CreateResponse struct {
	ShortUrl string
	LongUrl string
	ErrorMsg
}

func (s *service) Create(entry UrlEntry) (CreateResponse,error) {
	created,err := s.Repo.FindByLongurl(entry.LongUrl)
	if(err != nil) {
		fmt.Println(err.Error())
	}
	if(created.LongUrl == "") {
		created,err = s.Repo.Create(entry)
		if(err != nil) {
			return CreateResponse{}, err
		}
	}
	return CreateResponse{LongUrl:created.LongUrl, ShortUrl:created.ShortUrl},nil
}

func (s *service) Get(shortUrl string) (CreateResponse,error) {
	entry,err := s.Repo.Get(shortUrl)
	if(err != nil || entry.ShortUrl == "") {
		return CreateResponse{},utils.ErrNotFound
	}
	return CreateResponse{LongUrl:entry.LongUrl, ShortUrl:entry.ShortUrl},nil
}


//NewService is abstract service implementation for db calls
func NewService(repo Repository, redisClient utils.RedisClient) Service {
	return &service{Repo: repo, RedisClient:redisClient}
}

