package cacheredis

import (
	"my-fiber-app/mod/pkg/utils"

	"github.com/go-redis/redis/v8"
)

var (
	rdb  *redis.Client
)

// InitRedis initializes the Redis client using environment variables.
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     utils.GetEnv("REDIS_ADDR"),
		Password: utils.GetEnv("REDIS_PASSWORD"),
		DB:       0,
	})
}

// GetRedisClient returns the Redis client instance, initializing it if necessary.
func GetRedisClient() *redis.Client {
	if rdb == nil {
		InitRedis()
	}
	return rdb
}