package helper

import "github.com/go-redis/redis/v8"

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "172.27.0.2:6379",
	Username: "fariz",
	Password: "rahasia",
})
