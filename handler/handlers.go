package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishangordhan/go-url-shortener/shortener"
	"github.com/krishangordhan/go-url-shortener/store"
)

type URLCreationRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserID  string `json:"user_id"  binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	var creationRequest URLCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := shortener.GenerateShortLink(creationRequest.LongURL, creationRequest.UserID)
	store.SaveURLMapping(shortURL, creationRequest.LongURL, creationRequest.UserID)

	host := "http://localhost:9808/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "Short URL created successfully",
		"short_url": host + shortURL,
	})
}

func HandleShortURLRedirect(c *gin.Context) {
	shortURL := c.Param("shortURL")
	initialURL := store.RetrieveURL(shortURL)
	c.Redirect(http.StatusFound, initialURL)
}
