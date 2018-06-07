package utils

import (
	"gopkg.in/redis.v4"
	"time"
)

type redisClient struct {
	client *redis.Client
}

type RedisClient interface {
	SetValueInCache(key string, value string, time time.Duration) error
	GetValueFromCache(key string) (string, error)
}

func (r *redisClient) SetValueInCache(key string, value string, time time.Duration) error {
	return r.client.Set(key, value, time).Err()
}

func (r *redisClient) GetValueFromCache(key string) (string, error) {
	return r.client.Get(key).String(), nil
}

func NewRedisClient() RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	return &redisClient{client:client}
}
