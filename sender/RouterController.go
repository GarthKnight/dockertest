package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"strconv"
	"time"
)

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func InitRouter(rdb *redis.Client) {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/timestamp", func(context *gin.Context) {
		t := strconv.FormatInt(time.Now().Unix(), 10)
		err := rdb.Set("timestamp", t, 0).Err()
		if err != nil {
			fmt.Println(err)
		}

	})

	dataMap := make(map[string]interface{})
	dataMap["time"] = time.Now().Unix()

	router.POST("/redis/data", func(c *gin.Context) {
		data, err := rdb.XAdd(&redis.XAddArgs{
			Stream: streamName,
			ID:     "*",
			Values: dataMap,
		}).Result()

		fmt.Println("[REDIS_DATA_BASE] " + data)

		if err != nil {
			fmt.Println(err)
		}
	})

	runError := router.Run("0.0.0.0:8080")

	if runError != nil {
		fmt.Println("Router error")
	}
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
