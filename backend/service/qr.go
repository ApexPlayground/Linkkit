package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/model"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

type QRService struct {
	db       *gorm.DB
	clickSvc *ClickService
}

func NewQRService(db *gorm.DB, clickSvc *ClickService) *QRService {
	return &QRService{
		db:       db,
		clickSvc: clickSvc,
	}
}

func (s *QRService) CreateQRCode(userID uint, destinationURL string) (*model.QRCode, error) {
	qr := &model.QRCode{
		UserID:      userID,
		OriginalURL: destinationURL,
	}

	// Save first to get the ID
	if err := s.db.Create(qr).Error; err != nil {
		return nil, fmt.Errorf("failed to create QR code: %w", err)
	}

	// Generate image using the ID
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	qrURL := fmt.Sprintf("%s/qr/%d", baseURL, qr.ID)
	qrImage, err := qrcode.Encode(qrURL, qrcode.Medium, 400)
	if err != nil {
		// Rollback the created record since image failed
		s.db.Delete(qr)
		return nil, fmt.Errorf("failed to generate QR image: %w", err)
	}

	// Update only the image field
	if err := s.db.Model(qr).Update("qr_image", qrImage).Error; err != nil {
		s.db.Delete(qr)
		return nil, fmt.Errorf("failed to save QR image: %w", err)
	}

	// Cache in Redis
	cacheKey := fmt.Sprintf("qr:%d", qr.ID)
	payload, _ := json.Marshal(map[string]interface{}{
		"id":           qr.ID,
		"user_id":      qr.UserID,
		"original_url": qr.OriginalURL,
	})
	config.RDB.Set(config.Ctx, cacheKey, payload, 24*time.Hour)

	return qr, nil
}

// ResolveQR - Resolve QR code and track scan using existing click service
func (s *QRService) ResolveQR(
	qrID uint,
	ip string,
	userAgent string,
	referrer string,
) (string, error) {
	var destinationURL string

	// Try Redis cache first
	cacheKey := fmt.Sprintf("qr:%d", qrID)
	if val, err := config.RDB.Get(config.Ctx, cacheKey).Result(); err == nil {
		var cached map[string]interface{}
		if json.Unmarshal([]byte(val), &cached) == nil {
			if url, ok := cached["original_url"].(string); ok {
				destinationURL = url
			}
		}
	}

	// Cache miss, query DB
	if destinationURL == "" {
		var qr model.QRCode
		if err := s.db.Where("id = ?", qrID).First(&qr).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return "", gorm.ErrRecordNotFound
			}
			return "", err
		}

		destinationURL = qr.OriginalURL

		// Cache in Redis
		payload, _ := json.Marshal(map[string]interface{}{
			"id":           qr.ID,
			"user_id":      qr.UserID,
			"original_url": qr.OriginalURL,
		})
		config.RDB.Set(config.Ctx, cacheKey, payload, 24*time.Hour)
	}

	// Reuse existing click service to handles all the worker processing
	s.clickSvc.TrackQRScan(qrID, ip, userAgent, referrer)

	return destinationURL, nil
}

func (s *QRService) GetUserQRCodes(userID uint) ([]model.QRCode, error) {
	var qrs []model.QRCode
	// Fetch all QR codes for the user, newest first
	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&qrs).Error; err != nil {
		return nil, err
	}
	return qrs, nil
}

func (s *QRService) DeleteQRCode(qrID uint, userID uint) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Ensure user owns QR code before deleting
		result := tx.Where("id = ? AND user_id = ?", qrID, userID).Delete(&model.QRCode{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("QR code not found or access denied")
		}

		// Clean up all analytics related to user qr code
		for _, m := range []interface{}{
			&model.Click{},
			&model.DailyQRStats{},
			&model.QRCountryStats{},
			&model.QRDeviceStats{},
			&model.QRBrowserStats{},
			&model.QRReferrerStats{},
		} {
			if err := tx.Where("qr_id = ?", qrID).Delete(m).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	// clear cache after DB transaction succeeds
	config.RDB.Del(config.Ctx, fmt.Sprintf("qr:%d", qrID))
	return nil
}
