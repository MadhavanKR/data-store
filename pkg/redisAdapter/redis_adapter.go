package redisAdapter

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectAndStore() {
	redis.NewClient(&redis.Options{})
}
