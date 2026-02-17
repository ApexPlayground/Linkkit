// service/shortener.go
package service

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/model"
	"github.com/ApexPlayground/Linkkit/util"
	"gorm.io/gorm"
)

var (
	ErrURLEmpty     = errors.New("URL cannot be empty")
	ErrURLTooLong   = errors.New("URL too long")
	ErrURLInvalid   = errors.New("invalid URL")
	ErrLinkExists   = errors.New("link already exists")
	ErrLinkNotFound = errors.New("link not found or access denied")
)

type LinkService struct {
	db       *gorm.DB
	clickSvc *ClickService
}

func NewRedirectService(db *gorm.DB, clickSvc *ClickService) *LinkService {
	return &LinkService{
		db:       db,
		clickSvc: clickSvc,
	}
}

func CreateShortLink(userID uint, longUrl string) (*model.Link, error) {
	const (
		codeLength = 7
		maxRetries = 5
		maxURLLen  = 2048
	)

	longUrl = strings.TrimSpace(longUrl)

	if len(longUrl) == 0 {
		return nil, ErrURLEmpty
	}
	if len(longUrl) > maxURLLen {
		return nil, ErrURLTooLong
	}
	parsed, err := url.ParseRequestURI(longUrl)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return nil, ErrURLInvalid
	}

	// Check if user already shortened this URL
	var existing model.Link
	if err := config.DB.Where("user_id = ? AND long_url = ?", userID, longUrl).First(&existing).Error; err == nil {
		return &existing, ErrLinkExists
	}

	// Generate short code
	for range maxRetries {
		code, err := GenerateShortCode(codeLength)
		if err != nil {
			return nil, fmt.Errorf("failed to generate shortcode")
		}

		link := &model.Link{
			UserID:    userID,
			LongUrl:   longUrl,
			ShortCode: code,
		}

		err = config.DB.Create(&link).Error
		if err == nil {
			return link, nil
		}

		if strings.Contains(err.Error(), "duplicate key") {
			continue
		}

		return nil, fmt.Errorf("could not save link: %v", err)
	}

	return nil, fmt.Errorf("could not generate unique shortcode after %d retries", maxRetries)
}

func (s *LinkService) ResolveLink(
	shortcode string,
	ip string,
	userAgent string,
	referrer string,
) (string, error) {
	var linkID uint
	var longURL string

	// Try Redis cache first
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

	// Cache miss, query DB
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

	s.clickSvc.TrackClick(linkID, ip, userAgent, referrer)
	return longURL, nil
}

func GenerateShortCode(length int) (string, error) {
	const maxAttempts = 10
	for range maxAttempts {
		b := make([]byte, 8)
		_, err := rand.Read(b)
		if err != nil {
			return "", fmt.Errorf("failed to generate random bytes: %v", err)
		}
		randNum := binary.BigEndian.Uint64(b)
		encoded := util.Base62Encode(int(randNum % (1 << 62)))

		// fix for annoying empty encoded value
		if encoded == "" {
			continue
		}
		if len(encoded) >= length {
			return encoded[:length], nil
		}
	}
	return "", fmt.Errorf("failed to generate shortcode after %d attempts", maxAttempts)
}

func (s *LinkService) GetUserShortLinks(userID uint) ([]model.Link, error) {
	var links []model.Link

	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

func (s *LinkService) DeleteLink(linkID uint, userID uint) error {
	// Fetch shortcode before deleting for redis
	var link model.Link
	if err := s.db.Where("id = ? AND user_id = ?", linkID, userID).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("link not found or access denied")
		}
		return err
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Delete the link
		if err := tx.Delete(&link).Error; err != nil {
			return err
		}

		// Clean up all analytics related to this link
		for _, m := range []interface{}{
			&model.Click{},
			&model.DailyLinkStats{},
			&model.LinkCountryStats{},
			&model.LinkDeviceStats{},
			&model.LinkBrowserStats{},
			&model.LinkReferrerStats{},
		} {
			if err := tx.Where("link_id = ?", linkID).Delete(m).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	config.RDB.Del(config.Ctx, link.ShortCode)
	return nil
}
