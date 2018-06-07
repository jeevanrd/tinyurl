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
	"io/ioutil"
	"encoding/json"
	"strconv"
)

const (
	defaultPort = 8090
)
type Config struct {
	ApiPort string	`json:"api_port"`
	DbUri	string 	`json:"db_uri"`
	DbName	string	`json:"db_name"`
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := ioutil.ReadFile(pwd + "/config.json")
	if err != nil {
		fmt.Println("Error opening JSON file:")
		os.Exit(1)
	}

	var config Config
	err = json.Unmarshal(jsonFile, &config)

	if err != nil {
		fmt.Print("Error:", err)
	}

	dbObj,err := utils.Connect(config.DbUri, config.DbName)
	if(err !=nil) {
		log.Debug("Mongo error","unable to connect to mongo",err.Error())
		os.Exit(0)
	}

	r := mux.NewRouter()
	ctx := context.Background()

	repo := shorturl.NewRepository(dbObj)
	rClient := utils.NewRedisClient()
	sService := shorturl.NewService(repo, rClient)

	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(os.Stdout)

	router := shorturl.MakeHandler(r, logger, ctx, sService)
	if(config.ApiPort == "") {
		config.ApiPort = strconv.Itoa(defaultPort)
	}
	http.ListenAndServe(":"+config.ApiPort, router)
}
