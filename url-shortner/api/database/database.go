package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redisClient {
	rdb := NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ARRD"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}