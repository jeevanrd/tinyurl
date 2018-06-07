# tinyurl
Service to generate tinyurl

## Prerequisites:
- Install golang
- Install MongoDb
- Set GOPATH and GOROOT environment variables

## How to install dependencies
Install godep tool
```
go get github.com/tools/godep
```
To install all project dependencies
```
godep restore
```
### How to start service?
go run main.go

### How to run unit tests?
ginkgo -r

### APIs Available
- GET API

```
curl -X GET \
  http://localhost:3000/tinyurl/5b18840edf87ce1b5bc3c3d0 \
  -H 'content-type: application/json' \
  -d '{
	"longUrl": "http://dubdub2.com"
}'
```

- POST API

```
curl -X POST \
  http://localhost:3000/tinyurl \
  -H 'content-type: application/json' \
  -d '{
	"longUrl": "http://dubdub3326243.com"
}'
```
