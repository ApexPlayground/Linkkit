package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Redirect(db *gorm.DB, c *gin.Context) {
	shortcode := c.Param("shortcode")

	// Check Redis first
	longURL, err := config.RDB.Get(config.Ctx, shortcode).Result()
	if err == nil {
		c.Redirect(http.StatusMovedPermanently, longURL)
		return
	}

	var link model.Link
	if err := db.Where("short_code = ?", shortcode).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	if err := config.RDB.Set(config.Ctx, shortcode, link.LongUrl, 24*time.Hour).Err(); err != nil {
		fmt.Println("Redis set error:", err)
	}

	c.Redirect(http.StatusMovedPermanently, link.LongUrl)
}
