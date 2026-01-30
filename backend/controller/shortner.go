package controller

import (
	"log"
	"net/http"

	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
)

func ShortnerController(c *gin.Context) {
	var body struct {
		LongUrl string `json:"long_url"`
	}

	// read request
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Failed to parse shortener request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// call service
	link, err := service.CreateShortLink(body.LongUrl)
	if err != nil {
		log.Println("Failed to create short link:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	shortUrl := FormatShortURL(link.ShortCode)

	// success response
	c.JSON(http.StatusOK, gin.H{
		"long_url":  link.LongUrl,
		"short_url": shortUrl,
	})
}

const ShortURLPrefix = "http://127.0.0.1:8080/"

func FormatShortURL(shortCode string) string {
	return ShortURLPrefix + shortCode
}
