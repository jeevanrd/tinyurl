// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import shorturl "tiny_url/shorturl"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: entry
func (_m *Service) Create(entry shorturl.UrlEntry) (shorturl.CreateResponse, error) {
	ret := _m.Called(entry)

	var r0 shorturl.CreateResponse
	if rf, ok := ret.Get(0).(func(shorturl.UrlEntry) shorturl.CreateResponse); ok {
		r0 = rf(entry)
	} else {
		r0 = ret.Get(0).(shorturl.CreateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(shorturl.UrlEntry) error); ok {
		r1 = rf(entry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: tinyurl
func (_m *Service) Get(tinyurl string) (shorturl.CreateResponse, error) {
	ret := _m.Called(tinyurl)

	var r0 shorturl.CreateResponse
	if rf, ok := ret.Get(0).(func(string) shorturl.CreateResponse); ok {
		r0 = rf(tinyurl)
	} else {
		r0 = ret.Get(0).(shorturl.CreateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tinyurl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}