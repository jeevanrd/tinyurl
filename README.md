# tinyurl

Prerequisites:
1.Install golang
2.Install MongoDb
3.Set GOPATH and GOROOT environment variables

How to start service?
 go run main.go

How to run unit tests?
 ginkgo -r

APIs Available
1.GET API

curl -X GET \
  http://localhost:3000/tinyurl/5b18840edf87ce1b5bc3c3d0 \
  -H 'content-type: application/json' \
  -d '{
	"longUrl": "http://dubdub2.com"
}'

2. POST API

curl -X POST \
  http://localhost:3000/tinyurl \
  -H 'content-type: application/json' \
  -d '{
	"longUrl": "http://dubdub3326243.com"
}'