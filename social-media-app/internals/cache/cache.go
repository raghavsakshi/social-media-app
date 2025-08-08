package cache

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/redis/go-redis/v9"
)

var cache *redis.Client

func Client() *redis.Client {
	return cache
}

func Connect() {
	ctx := context.Background()
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
		cache = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: "",
			DB:       0,
		})
	} else {
		redisOptions, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatalf("Error parsing Redis URL: %v", err)
			
		}
		cache = redis.NewClient(redisOptions)
	}
	cmd := cache.Ping(ctx)
	if cmd.Err() != nil {
		fmt.Println("Error connecting to Redis:", cmd.Err())
		panic(cmd.Err())
	}
	
	fmt.Println("Successfully connected to Redis:")
}

