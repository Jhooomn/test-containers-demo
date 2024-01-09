package cache

import "github.com/redis/go-redis/v9"

func SetUp(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr + ":6379", // "172.17.0.3:6379",
		Password: "",             // no password set
		DB:       0,              // use default DB
	})
}
