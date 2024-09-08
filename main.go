package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/krishangordhan/go-url-shortener/handler"
	"github.com/krishangordhan/go-url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey!~",
		})
	})

	r.POST("create", func(c *gin.Context) {
		handler.CreateShortURL(c)
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		handler.HandleShortURLRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
