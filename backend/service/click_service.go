package service

import (
	"time"

	"github.com/ApexPlayground/Linkkit/model"
	"gorm.io/gorm"
)

type ClickService struct {
	DB *gorm.DB
}

func NewClickService(db *gorm.DB) *ClickService {
	return &ClickService{DB: db}
}

func (s *ClickService) TrackClick(linkID uint, ip, userAgent, referrer string) {
	go func() {
		click := model.Click{
			LinkID:    linkID,
			CreatedAt: time.Now(),
			IP:        ip,
			UserAgent: userAgent,
			Referrer:  referrer,
		}

		if err := s.DB.Create(&click).Error; err != nil {

			// fmt.Println("Click tracking error:", err)
		}

	}()
}
