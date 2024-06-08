package connection

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(addr, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func PingRedis(client *redis.Client) error {
	ctx := context.Background()
	return client.Ping(ctx).Err()
}
