package shorturl

import (
	"github.com/gorilla/mux"
	"context"
	"net/http"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"tiny_url/utils"
	"encoding/json"
	"strings"
	"fmt"
)

type GetUrlRequest struct {
	ShortUrl string
}

type CreateUrlRequest struct {
	LongUrl	string
}

func MakeHandler(r *mux.Router, logger kitlog.Logger, ctx context.Context, s Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createTinyUrlHandler := kithttp.NewServer(
		makeCreateEndpoint(s),
		decodeCreateRequest,
		encodeResponse,
		opts...,
	)

	getTinyUrlHandler := kithttp.NewServer(
		makeGetEndpoint(s),
		decodeGetRequest,
		encodeResponse,
		opts...,
	)

	r.Handle("/tinyurl/{shorturl}", getTinyUrlHandler).Methods("GET")
	r.Handle("/tinyurl", createTinyUrlHandler).Methods("POST")
	return r
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error)  {
	vars := mux.Vars(r)
	shorturl, ok := vars["shorturl"]
	if !ok {
		return nil, utils.ErrBadRoute
	}

	if(strings.Trim(shorturl, "") == "") {
		return nil, utils.ErrInvalidArgument
	}
	return GetUrlRequest {
		ShortUrl:shorturl,
	}, nil
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error)  {
	var url UrlEntry
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		return nil, err
	}
	if(strings.Trim(url.LongUrl, "") == "") {
		return nil, utils.ErrInvalidArgument
	}
	return url, nil
}


func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println(21)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case utils.ErrInvalidArgument,utils.JsonParseError,utils.ErrBadRoute:
		w.WriteHeader(http.StatusBadRequest)
	case utils.ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println(31)
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}