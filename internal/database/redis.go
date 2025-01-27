package database

import (
	"context"
	"fmt"
	"product-service/internal/config"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func InitRedis(c *config.Redis) (*redis.Client, error) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
		Password: c.Password, // No password set
		DB:       0,          // Use default DB
	})

	// Test Redis connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	return RedisClient, nil
}
