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

type redisLink struct {
	ID      uint   `json:"id"`
	LongURL string `json:"long_url"`
}

func (s *RedirectService) Resolve(
	shortcode string,
	ip string,
	ua string,
	referer string,
) (string, error) {

	// 1. Try Redis cache
	val, err := config.RDB.Get(config.Ctx, shortcode).Result()
	if err == nil {
		var cached redisLink
		if json.Unmarshal([]byte(val), &cached) == nil {
			s.clickSvc.TrackClick(cached.ID, ip, ua, referer)
			return cached.LongURL, nil
		}
	}

	// 2. Fetch from DB
	var link model.Link
	if err := s.db.Where("short_code = ?", shortcode).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", gorm.ErrRecordNotFound
		}
		return "", err
	}

	// 3. Cache in Redis
	payload, _ := json.Marshal(redisLink{
		ID:      link.ID,
		LongURL: link.LongUrl,
	})
	config.RDB.Set(config.Ctx, shortcode, payload, 24*time.Hour)

	// 4. Track click
	s.clickSvc.TrackClick(link.ID, ip, ua, referer)

	return link.LongUrl, nil
}
