package conf

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var redisClient *redis.Client
var expiredTime = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct{}

func InitRedis() (*RedisClient, error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any) error {
	return redisClient.Set(context.Background(), key, value, expiredTime).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return redisClient.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return redisClient.Del(context.Background(), key...).Err()
}
