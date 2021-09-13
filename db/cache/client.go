package cache

import (
	"github.com/go-redis/redis/v8"
	"web-svc/config"
)

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: "",
		DB:       0,
	})

	return client
}
