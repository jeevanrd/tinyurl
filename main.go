package main

import (
	"github.com/gorilla/mux"
	"context"
	"tiny_url/shorturl"
	"tiny_url/utils"
	"net/http"
	"github.com/prometheus/common/log"
	kitlog "github.com/go-kit/kit/log"
	"fmt"
	"os"
)

func main() {
	uri := "mongodb://localhost:27017/"
	dbName := "urls"
	port := "3000"

	fmt.Println(1)
	dbObj,err := utils.Connect(uri, dbName)
	fmt.Println(1)
	if(err !=nil) {
		log.Debug("Mongo error","unable to connect to mongo",err.Error())
	}

	fmt.Println(1)
	r := mux.NewRouter()
	ctx := context.Background()
	fmt.Println(2)

	repo := shorturl.NewRepository(dbObj)
	rClient := utils.NewRedisClient()
	fmt.Println(3)
	sService := shorturl.NewService(repo, rClient)


	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(os.Stdout)

	router := shorturl.MakeHandler(r, logger, ctx, sService)
	http.ListenAndServe(":"+port, router)
}
