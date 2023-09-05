package databases

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func StartRedis(host string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})

	// check connexion
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("Unable to start redis")
	}
	// assign client to global variable (possible by dependency injection too)
	RedisClient = client
	return client
}
