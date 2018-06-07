package shorturl_test

import (
	"net/http"
	"context"
	"net/http/httptest"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"bytes"
	"tiny_url/shorturl"
	"tiny_url/shorturl/mocks"
	"io"

	"github.com/go-kit/kit/log"
	"os"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"fmt"
)

func addUserInfo(request *http.Request) *http.Request {
	ctx := request.Context()
	ctx = context.WithValue(ctx, "request-method", request.Method)
	request = request.WithContext(ctx)
	return request
}

func addHTTPMethod(r *http.Request, ctx context.Context, httpMethod string) *http.Request {
	ctx = context.WithValue(ctx, "request-method", httpMethod)
	return r.WithContext(ctx)
}

func PrepareGetRequest(url string) *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	request = addUserInfo(request)
	return addHTTPMethod(request, request.Context(), "GET")
}

func PreparePostRequest(url string, body io.Reader) *http.Request {
	request, _ := http.NewRequest("POST", url, body)
	request = addUserInfo(request)
	return addHTTPMethod(request, request.Context(), "POST")
}


var _ = Describe("Transport", func() {

	var id1 = bson.NewObjectId()
	var url1 = shorturl.UrlEntry{ID:id1, ShortUrl:id1.Hex(), LongUrl:"http://www.google.com"}
	//var url2 = shorturl.UrlEntry{ID:id2, ShortUrl:id2.Hex(), LongUrl:"http://www.google1.com"}


	var recorder *httptest.ResponseRecorder
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	mockService := &mocks.Service{}
	handler := shorturl.MakeHandler(mux.NewRouter(),logger , context.Background(), mockService)
	mockService.On("Create", url1).Return(shorturl.CreateResponse{ShortUrl:url1.ShortUrl, LongUrl:url1.LongUrl}, nil)
	mockService.On("Get", url1.ShortUrl).Return(shorturl.CreateResponse{ShortUrl:url1.ShortUrl, LongUrl:url1.LongUrl}, nil)


	BeforeEach(func() {
		// Set up a new server before each test and record HTTP responses.
		recorder = httptest.NewRecorder()
	})

	Describe("GET /tinyurl/:id", func() {
		Context("get tinyurl", func() {
			url := "/tinyurl/"+url1.ShortUrl
			request := PrepareGetRequest(url)

			It("returns 200 OK", func() {
				handler.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})
			It("returns id details in result", func() {
				handler.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).Should(ContainSubstring(url1.ShortUrl))
			})
		})
	})


	Describe("POST /tinyurl", func() {
		Context("POST tinyurl", func() {
			url := "/tinyurl"
			byteData,_ := json.Marshal(url1)
			fmt.Println(string(byteData))
			request := PreparePostRequest(url, bytes.NewBuffer(byteData))

			It("create tinyurl", func() {
				handler.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})

		Context("POST tinyurl", func() {
			url := "/tinyurl"
			request := PreparePostRequest(url, bytes.NewBuffer([]byte("{}")))
			It("should get badrequest", func() {
				handler.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})

		})
	})
})


