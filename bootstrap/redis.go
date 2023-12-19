package bootstrap

import (
	"context"

	"srating/utils"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(env *Env) *redis.Client {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisURL,
		Password: env.RedisPassword,
		DB:       env.RedisDB,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		utils.LogFatal(err, "Failed to connect to Redis")
		return nil
	}

	return client
}

func CloseRedisConnection(client *redis.Client) {
	err := client.Close()
	if err != nil {
		utils.LogFatal(err, "Failed to close Redis connection")
		return
	}
}
