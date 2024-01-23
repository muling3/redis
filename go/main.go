package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	val2 := redisClient.SetEx(ctx, "yearDay", time.Now().YearDay(), time.Hour * 1)
	log.Println("Result cmd", val2.FullName())

	allKeys, err := redisClient.MGet(ctx, "appName", "yearDay").Result()
	if err != nil {
		log.Println("Error in getting keys")
	}

	// printing the keys
	for k, v := range allKeys {
		log.Print(k, " - ", v)
	}

	// get all keys in the database
	keys, err := redisClient.Keys(ctx, "*").Result()
	if err != nil {
		log.Println("Error in getting keys")
	}

	// printing the keys
	for k, v := range keys {
		log.Print(k, " - ", v)
	}
	log.Println("Key Value Retrieved is --> ", val)
}
