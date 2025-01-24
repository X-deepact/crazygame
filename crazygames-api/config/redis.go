package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     AppConfig.RedisHost + ":" + AppConfig.RedisPort,
		Password: AppConfig.RedisPassword,
		DB:       0,
	})

	ctx := context.Background()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return client
}
