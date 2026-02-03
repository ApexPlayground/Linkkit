package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/model"
	"gorm.io/gorm"
)

type RedirectService struct {
	db       *gorm.DB
	clickSvc *ClickService
}

func NewRedirectService(db *gorm.DB, clickSvc *ClickService) *RedirectService {
	return &RedirectService{
		db:       db,
		clickSvc: clickSvc,
	}
}

func (s *RedirectService) Resolve(
	shortcode string,
	ip string,
	userAgent string,
	referrer string,
) (string, error) {

	// Try Redis cache first
	var linkID uint
	var longURL string

	if val, err := config.RDB.Get(config.Ctx, shortcode).Result(); err == nil {
		var cached struct {
			ID      uint   `json:"id"`
			LongURL string `json:"long_url"`
		}
		if json.Unmarshal([]byte(val), &cached) == nil {
			linkID = cached.ID
			longURL = cached.LongURL
		}
	}

	// Cache miss → query DB
	if linkID == 0 {
		var link model.Link
		if err := s.db.Where("short_code = ?", shortcode).First(&link).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return "", gorm.ErrRecordNotFound
			}
			return "", err
		}

		linkID = link.ID
		longURL = link.LongUrl

		// Cache in Redis
		payload, _ := json.Marshal(struct {
			ID      uint   `json:"id"`
			LongURL string `json:"long_url"`
		}{linkID, longURL})

		config.RDB.Set(config.Ctx, shortcode, payload, 24*time.Hour)
	}

	// Track click asynchronously
	s.clickSvc.TrackClick(linkID, ip, userAgent, referrer)

	return longURL, nil
}
