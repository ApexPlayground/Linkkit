package controller

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var qrSvc *service.QRService

func InitQRController(svc *service.QRService) {
	qrSvc = svc
}

func CreateQRCode(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req struct {
		OriginalURL string `json:"original_url" binding:"required,url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qr, err := qrSvc.CreateQRCode(userID, req.OriginalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create QR"})
		return
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":           qr.ID,
		"original_url": qr.OriginalURL,
		"qr_url":       baseURL + "/qr/" + strconv.FormatUint(uint64(qr.ID), 10),
		"created_at":   qr.CreatedAt,
	})
}

func QRRedirect(c *gin.Context) {
	qrID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid QR code ID"})
		return
	}

	destinationURL, err := qrSvc.ResolveQR(
		uint(qrID),
		c.ClientIP(),
		c.Request.UserAgent(),
		c.Request.Referer(),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "QR code not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.Redirect(http.StatusFound, destinationURL)
}

func QRListController(c *gin.Context) {
	userID := c.GetUint("user_id")

	qrs, err := qrSvc.GetUserQRCodes(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch QR codes"})
		return
	}

	response := make([]map[string]interface{}, 0)
	for _, qr := range qrs {
		response = append(response, map[string]interface{}{
			"id":           qr.ID,
			"original_url": qr.OriginalURL,
			"qr_url":       "/qr/" + strconv.FormatUint(uint64(qr.ID), 10),
			"created_at":   qr.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

func QRDeleteController(c *gin.Context) {
	userID := c.GetUint("user_id")

	qrID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid QR code ID"})
		return
	}

	if err := qrSvc.DeleteQRCode(uint(qrID), userID); err != nil {
		if err.Error() == "QR code not found or access denied" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete QR code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "QR code deleted successfully",
	})
}
