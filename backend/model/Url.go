package model

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID        uint           `gorm:"primaryKey"`
	LongUrl   string         `gorm:"not null"`
	ShortCode string         `gorm:"uniqueIndex"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Click struct {
	ID        uint      `gorm:"primaryKey"`
	LinkID    uint      `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Link reference
	CreatedAt time.Time `gorm:"not null"`                                                     // Timestamp of click
	IP        string    `gorm:"not null"`                                                     // User IP
	UserAgent string    `gorm:"not null"`                                                     // Browser/device info
	Referrer  string
	Country   string         `gorm:"default:null"`
	Device    string         `gorm:"default:null"`
	Browser   string         `gorm:"default:null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
