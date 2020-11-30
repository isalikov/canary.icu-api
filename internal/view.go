package internal

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// ViewPost response with single post
func ViewPost(ctx context.Context, rdb *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		post := extractPost(ctx, c.Param("ID"), rdb)

		c.HTML(http.StatusOK, "view.tmpl", post)

	}
}
