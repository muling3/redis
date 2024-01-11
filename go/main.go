package main

import (
	"context"
	"fmt"
	"log"

	"github.com/muling3/redis/go/config"
)

func main() {
	log.Println("Starting Redis GO 🐇🐇")
	ctx := context.Background()

	// connect to redis
	redisClient := config.RedisConn()

	// create a var
	if err := redisClient.Set(ctx, "appName", "Redis Go", 0).Err(); err != nil {
		log.Println("Error in creating key")
	}

	fmt.Println("Key set successfully")

	// getting the key
	val, err := redisClient.Get(ctx, "appName").Result()
	if err != nil {
		log.Println("Error in getting key")
	}

	log.Println("Key Value Retrieved is --> ", val)
}
