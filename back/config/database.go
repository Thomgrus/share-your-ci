package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

const (
	port     = 6379
	password = ""
	db       = 0
)

var client *redis.Client

// DatabaseInit : initialize database connection
func DatabaseInit() {
	var redisAddr = fmt.Sprintf("%s:%d", os.Getenv("REDIS_HOSTNAME"), port)

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}
}

// Db : Getter for db var
func Client() *redis.Client {
	return client
}
