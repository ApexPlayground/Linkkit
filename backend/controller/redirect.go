package controller

import (
	"encoding/json"
	"errors"

	"net/http"
	"time"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/model"
	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RedisLink struct {
	ID      uint   `json:"id"`
	LongURL string `json:"long_url"`
}

func Redirect(db *gorm.DB, clickSvc *service.ClickService, c *gin.Context) {
	shortcode := c.Param("shortcode")
	var link model.Link

	// Try Redis cache first
	val, err := config.RDB.Get(config.Ctx, shortcode).Result()
	if err == nil {
		var cached RedisLink
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			clickSvc.TrackClick(cached.ID, c.ClientIP(), c.Request.UserAgent(), c.Request.Referer())
			c.Redirect(302, cached.LongURL)
			return
		}
	}

	// Fetch from DB
	if err := db.Where("short_code = ?", shortcode).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Save to Redis
	cached := RedisLink{ID: link.ID, LongURL: link.LongUrl}
	if data, err := json.Marshal(cached); err == nil {
		config.RDB.Set(config.Ctx, shortcode, data, 24*time.Hour)
	}

	// Track click
	clickSvc.TrackClick(link.ID, c.ClientIP(), c.Request.UserAgent(), c.Request.Referer())

	c.Redirect(302, link.LongUrl)
}
