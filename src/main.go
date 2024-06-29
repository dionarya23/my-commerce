package main

import (
	"fmt"

	"dionpamungkas.com/my-commerce/src/drivers/db"
	"dionpamungkas.com/my-commerce/src/drivers/redis"
	"dionpamungkas.com/my-commerce/src/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dbConnection, err := db.CreateConnection()
	redisConnection, errRedis := redis.CreateConnection()

	if err != nil {
		fmt.Println("Error creating database connection:", err)
		return
	}

	if errRedis != nil {
		fmt.Println("Error creating redis connection:", err)
		return
	}

	defer func() {
		if err := dbConnection.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	h := http.New(&http.Http{
		DB:          dbConnection,
		RedisClient: redisConnection,
	})

	h.Launch()
}
