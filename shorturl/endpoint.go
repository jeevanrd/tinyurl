package shorturl

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)


type Endpoint struct {
	s Service
}

func makeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UrlEntry)
		c, err := s.Create(req)
		return c, err
	}
}

func makeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUrlRequest)
		c, err := s.Get(req.ShortUrl)
		return c, err
	}
}