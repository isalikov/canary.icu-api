package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/isalikov/canary.icu-api/internal"
)

var ctx = context.Background()

func getConnectionConfig() redis.Options {
	host := os.Getenv("REDIS_HOST")

	if host == "" {
		panic("Can't get Redist host")
	}

	port := os.Getenv("REDIS_PORT")

	if port == "" {
		panic("Can't get Redist port")
	}

	addr := fmt.Sprintf("%v:%v", host, port)

	return redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	}
}

func main() {
	config := getConnectionConfig()
	rdb := redis.NewClient(&config)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/j", internal.GetFeed(ctx, rdb))
	r.GET("/j/:ID", internal.GetPost(ctx, rdb))
	r.GET("/v/:ID", internal.ViewPost(ctx, rdb))

	r.Run()
}
