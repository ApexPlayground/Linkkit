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
	LinkID    uint      `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"not null"`
	IP        string    `gorm:"not null"`
	UserAgent string    `gorm:"not null"`
	Referrer  string
}
