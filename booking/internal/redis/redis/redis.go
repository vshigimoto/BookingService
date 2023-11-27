package redis

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

func NewRedisClient(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return rdb
}

func GetValue(rdb *redis.Client, key string) string {
	val, err := rdb.Get(key).Result()
	if err == redis.Nil {
		log.Printf("key does not exist")
		return ""
	}
	if err != nil {
		return ""
	}
	return val
}

func SetValue(rdb *redis.Client, key, value string) error {
	exp := 30 * (time.Minute)
	err := rdb.Set(key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}
