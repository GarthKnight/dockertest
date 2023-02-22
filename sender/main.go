package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

const streamName = "datastream"

func main() {
	client := getRedisClient()

	go listenToRedisStream(client)

	InitRouter(client)
}

func getRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func listenToRedisStream(rdb *redis.Client) {

	id := "$"
	for {
		data, err := rdb.XRead(&redis.XReadArgs{
			Streams: []string{streamName, id},
			Block:   0,
		}).Result()

		if err != nil {
			fmt.Println(err)
		}

		for _, result := range data {
			for _, message := range result.Messages {
				id = message.ID
				fmt.Println("[REDIS_DATA_BASE]")
				fmt.Println(message)
			}
		}
	}

}
