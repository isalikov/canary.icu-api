package internal

import (
	"context"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type payload struct {
	ID    string
	Src   string
	Text  template.HTML
	Title string
}

// ViewPost response with single post
func ViewPost(ctx context.Context, rdb *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		post := extractPost(ctx, c.Param("ID"), rdb)

		payload := payload{
			ID:    post.ID,
			Src:   post.PreviewImage,
			Text:  template.HTML(post.Message),
			Title: post.Title,
		}

		c.HTML(http.StatusOK, "view.tmpl", payload)

	}
}
