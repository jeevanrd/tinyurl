package shorturl_test

import (
	//"context"
	"os"
	dbMocks "tiny_url/shorturl/mocks"
	"github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	utilMocks "tiny_url/utils/mocks"
	"tiny_url/shorturl"
	"gopkg.in/mgo.v2/bson"
	"tiny_url/utils"
	//"net/url"
)

var id1 = bson.NewObjectId()
var id2 = bson.NewObjectId()
var id3 = bson.NewObjectId()

var url1 = shorturl.UrlEntry{ID:id1, ShortUrl:id1.Hex(), LongUrl:"http://www.google.com"}
var url2 = shorturl.UrlEntry{ID:id2, ShortUrl:id2.Hex(), LongUrl:"http://www.google1.com"}
var url3 = shorturl.UrlEntry{ID:id3, ShortUrl:id3.Hex(), LongUrl:"http://www.google2.com"}


var _ = Describe("Service", func() {

	//ctx := context.Background()
	mockDatabaseRepo := &dbMocks.Repository{}
	mockRedisClient := &utilMocks.RedisClient{}

	mockDatabaseRepo.On("Get", url1.ShortUrl).Return(url1, nil)
	mockDatabaseRepo.On("Get", url2.ShortUrl).Return(url2, nil)
	mockDatabaseRepo.On("Get", url3.ShortUrl).Return(url2, utils.ErrNotFound)


	mockDatabaseRepo.On("Create", url2).Return(url2, nil)
	mockDatabaseRepo.On("Create", url1).Return(url1, nil)

	mockDatabaseRepo.On("FindByLongurl", url2.LongUrl).Return(shorturl.UrlEntry{}, nil)


	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	//	logger = &serializedLogger{Logger: logger}
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var s shorturl.Service
	s = shorturl.NewService(mockDatabaseRepo,mockRedisClient)

	Describe("GET /tiny_url/{id}", func() {
		Context("get longurl for shorturl", func() {

			It("error should be nil", func() {
				urlEntry, err := s.Get(url1.ShortUrl)
				Expect(err).To(BeNil())
				Expect(urlEntry.ShortUrl).To(Equal(url1.ShortUrl))
				Expect(urlEntry.LongUrl).To(Equal(url1.LongUrl))
			})

			It("should return not found nil", func() {
				_,err := s.Get(id3.Hex())
				Expect(err).To(Equal(utils.ErrNotFound))
			})
		})
	})

	Describe("GET /tiny_url/", func() {
		Context("create shorturl for longurl", func() {
			created, err := s.Create(url2)
			It("error should be nil", func() {
				Expect(err).To(BeNil())
				Expect(created.ShortUrl).To(Equal(url2.ShortUrl))
				Expect(created.LongUrl).To(Equal(url2.LongUrl))
			})
		})
	})
})
