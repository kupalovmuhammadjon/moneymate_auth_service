package redis

import (
	"auth_service/configs"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *configs.Config) (*redis.Client, error) {

	redisClient := *redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDBNumber,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &redisClient, nil
}
