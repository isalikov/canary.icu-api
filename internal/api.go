package internal

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/isalikov/canary.icu-api/internal/db"
)

func extractPost(ctx context.Context, key string, rdb *redis.Client) *db.Post {
	post := &db.Post{}
	v, _ := rdb.Get(ctx, key).Result()

	json.Unmarshal([]byte(v), post)

	return post
}

// GetPost response with single post
func GetPost(ctx context.Context, rdb *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		post := extractPost(ctx, c.Param("ID"), rdb)

		c.JSON(200, post)

	}
}

// GetFeed response with news feed
func GetFeed(ctx context.Context, rdb *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		v, _ := rdb.Keys(ctx, "*").Result()

		result := []db.Post{}

		for _, key := range v {
			post := extractPost(ctx, key, rdb)

			if post.ID != "" {
				result = append(result, *post)
			}
		}

		sort.SliceStable(result, func(a, b int) bool {
			return result[a].Timestamp < result[b].Timestamp
		})

		c.JSON(200, result)

	}
}
