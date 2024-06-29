package redis

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func CreateConnection() (*redis.Client, error) {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	portNumberRedis, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", os.Getenv("REDIS_HOST"), portNumberRedis),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	})

	return redisClient, nil
}
