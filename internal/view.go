package internal

import (
	"context"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type payload struct {
	Title string
	Src   string
	Text  template.HTML
}

// ViewPost response with single post
func ViewPost(ctx context.Context, rdb *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		post := extractPost(ctx, c.Param("ID"), rdb)

		payload := payload{
			Title: post.Title,
			Src:   post.PreviewImage,
			Text:  template.HTML(post.Message),
		}

		c.HTML(http.StatusOK, "view.tmpl", payload)

	}
}
