package config

import "github.com/redis/go-redis/v9"

func RedisConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  
	})

	return client
}
