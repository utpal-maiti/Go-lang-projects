package cacheredis

import (
	"my-fiber-app/mod/pkg/utils"

	"github.com/go-redis/redis/v8"
)

var (
	rdb  *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     utils.GetEnv("REDIS_ADDR"),
		Password: utils.GetEnv("REDIS_PASSWORD"),
		DB:       0,
	})
}
func GetRedisClient() *redis.Client {
	if rdb == nil {
		InitRedis()
	}
	return rdb
}	