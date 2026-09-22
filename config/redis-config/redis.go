package config

import (
	"fmt"
	"log"

	"github.com/Devrao-2006/Go-Limit/config"
	"github.com/redis/go-redis/v9"
)

const REDIS_DB = 2

var RedisClients map[int]*redis.Client = make(map[int]*redis.Client)

func GetClient(db int) *redis.Client {
	if RedisClients[db] != nil {
		return RedisClients[db]
	}

	client := redis.NewClient(&redis.Options{
		Addr: config.GetEnv("REDIS_URL"),
		DB:   db,
	})

	RedisClients[db] = client

	return client
}

func InitClients() error {
	client := GetClient(REDIS_DB)

	if client == nil {
		return fmt.Errorf("failed to initialize Redis")
	}

	fmt.Println("All the Redis Client Are Initialized")

	return nil
}

func CloseRedis() {
	for _, client := range RedisClients {
		if err := client.Close(); err != nil {
			log.Printf("Error Closing Redis: %v", err)
		}
	}
}
