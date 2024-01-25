package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/muling3/redis/go/config"
)

type Person struct {
	Name   string
	Gender string
}

func (p Person) String() string {
	return strings.Join([]string{p.Name, p.Gender}, " - ")
}

func main() {
	// myChan := make(chan int, 1)
	// myChan <- 12
	// v := <-myChan
	// log.Println("Value v ", v)

	log.Println("Starting Redis GO 🐇🐇")
	ctx := context.Background()

	// connect to redis
	redisClient := config.RedisConn()
	log.Println("Redis Client", redisClient.Conn())

	// create a var
	if err := redisClient.Set(ctx, "appName", "Redis Go", 0).Err(); err != nil {
		log.Println("Error in creating key", err)
	}
	fmt.Println("Key set successfully")

	// getting the key
	val, err := redisClient.Get(ctx, "appName").Result()
	if err != nil {
		log.Println("Error in getting key")
	}

	val2 := redisClient.SetEx(ctx, "yearDay", time.Now().YearDay(), time.Hour*1)
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

	// testing type assertions in go from interface type
	var myVal interface{}

	myVal = Person{
		Name:   "Mimi",
		Gender: "Female",
	}

	log.Printf("Type Casted Value is => %+v", myVal.(Person))

}
