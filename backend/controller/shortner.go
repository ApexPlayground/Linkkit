package controller

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var linkSvc *service.LinkService

func InitLinkController(svc *service.LinkService) {
	linkSvc = svc
}

var ShortURLPrefix = os.Getenv("BASE_URL")

func CreateShortLink(c *gin.Context) {
	userID := c.GetUint("user_id")

	var body struct {
		LongUrl string `json:"long_url"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Failed to parse shortener request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	link, err := service.CreateShortLink(userID, body.LongUrl)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrURLEmpty):
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL cannot be empty"})
			return
		case errors.Is(err, service.ErrURLTooLong):
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL is too long"})
			return
		case errors.Is(err, service.ErrURLInvalid):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		case errors.Is(err, service.ErrLinkExists):
			// Inform user link already exists
			c.JSON(http.StatusOK, gin.H{
				"long_url":    link.LongUrl,
				"short_url":   FormatShortURL(link.ShortCode),
				"message":     "Link already exists",
				"is_existing": true,
			})
			return
		default:
			log.Println("Failed to create short link:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}

	// Successfully created new link
	c.JSON(http.StatusCreated, gin.H{
		"long_url":    link.LongUrl,
		"short_url":   FormatShortURL(link.ShortCode),
		"is_existing": false,
	})
}

func LinkRedirect(c *gin.Context) {
	shortcode := c.Param("shortcode")

	longURL, err := linkSvc.ResolveLink(
		shortcode,
		c.ClientIP(),
		c.Request.UserAgent(),
		c.Request.Referer(),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.Redirect(http.StatusFound, longURL)
}

func FormatShortURL(shortCode string) string {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// Ensure baseURL ends with /
	if baseURL[len(baseURL)-1] != '/' {
		baseURL += "/"
	}

	return baseURL + shortCode
}

func LinkListController(c *gin.Context) {
	userID := c.GetUint("user_id")

	links, err := linkSvc.GetUserShortLinks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shortened links"})
		return
	}

	response := make([]map[string]interface{}, 0)
	for _, link := range links {
		response = append(response, map[string]interface{}{
			"id":         link.ID,
			"long_url":   link.LongUrl,
			"short_url":  ShortURLPrefix + link.ShortCode,
			"created_at": link.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

func LinkDeleteController(c *gin.Context) {
	userID := c.GetUint("user_id")

	linkID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link ID"})
		return
	}

	if err := linkSvc.DeleteLink(uint(linkID), userID); err != nil {
		if errors.Is(err, service.ErrLinkNotFound) {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete link"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Link deleted successfully",
	})
}
