package main

import (
	"fmt"

	"github.com/facufachin/go-url-shortener/handler"
	"github.com/facufachin/go-url-shortener/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenido a la api para el acortador de URLs",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
