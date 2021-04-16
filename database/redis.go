package database

import "github.com/go-redis/redis/v8"

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
}
