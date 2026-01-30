package controller

import (
	"net/http"

	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var redirectSvc *service.RedirectService

func InitRedirectController(svc *service.RedirectService) {
	redirectSvc = svc
}

func Redirect(c *gin.Context) {
	shortcode := c.Param("shortcode")

	longURL, err := redirectSvc.Resolve(
		shortcode,
		c.ClientIP(),
		c.Request.UserAgent(),
		c.Request.Referer(),
	)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.Redirect(http.StatusFound, longURL)
}
