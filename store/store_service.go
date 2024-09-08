package store

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type StorageService struct {
	RedisClient *redis.Client
}

var (
	storeService = &StorageService{}
)

const (
	CacheDuration = 6 * time.Hour
)

func InitializeStore() *StorageService {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing Redis: %v", err))
	}

	log.Printf("\nRedis initialized: %v", pong)
	storeService.RedisClient = RedisClient
	return storeService
}

func SaveURLMapping(shortURL string, originalURL string, userID string) {
	err := storeService.RedisClient.Set(shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortURL, originalURL))
	}
}

func RetrieveURL(shortURL string) string {
	result, err := storeService.RedisClient.Get(shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed retrieving url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	return result
}
