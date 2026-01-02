package controller

import (
	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
)

func ShortnerController(c *gin.Context) {
	var body struct {
		LongUrl string `json:"long_url"`
	}

	// read request
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	//call service and pass request
	link, err := service.CreateShortLink(body.LongUrl)
	if err != nil {
		serviceErr, isServiceErr := err.(service.ServiceError)
		if isServiceErr {
			c.JSON(serviceErr.Status, gin.H{"error": serviceErr.Message})
			return
		}

		// fallback generic error
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	shortUrl := FormatShortURL(link.ShortCode)

	//success response
	c.JSON(200, gin.H{
		"long_url":  link.LongUrl,
		"short_url": shortUrl,
	})
}

const ShortURLPrefix = "http://127.0.0.1:8080/"

func FormatShortURL(shortCode string) string {
	return ShortURLPrefix + shortCode
}
