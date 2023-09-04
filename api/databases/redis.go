package databases

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func StartRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Vérifier la connexion
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("Unable to start redis")
	}
	RedisClient = client
	return client
}
