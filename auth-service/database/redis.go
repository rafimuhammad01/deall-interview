package database

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"os"
)

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	return rdb
}
