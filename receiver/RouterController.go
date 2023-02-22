package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

func InitRouter(rdb *redis.Client) {
	router := gin.Default()
	router.GET("/timestamp", func(c *gin.Context) {
		val := getTimestampFromDB(rdb)
		c.IndentedJSON(http.StatusOK, val)
	})

	runError := router.Run("0.0.0.0:8080")
	if runError != nil {
		fmt.Println("Router error")
	}
}

func getTimestampFromDB(rdb *redis.Client) string {
	val, err := rdb.Get("timestamp").Result()

	if val == "" {
		val = "no timestamp"
	}

	if err != nil {
		fmt.Println("[REDIS] " + err.Error())
	}
	return val
}
