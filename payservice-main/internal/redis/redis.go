package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	ctx         context.Context
)

func Init(Address string) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: Address,
	})

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err := RedisClient.Set(ctx, "test", "gabdyq", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Redis client connected successfully...")

	return RedisClient
}
