package handler

import (
	"net/http"

	"github.com/facufachin/go-url-shortener/shortener"
	"github.com/facufachin/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "https://facu.gay/acortador/"
	c.JSON(200, gin.H{
		"message":   "URL Acortada Correctamente",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
