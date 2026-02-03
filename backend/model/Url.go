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

type ClickStat struct {
	ID        uint      `gorm:"primaryKey"`
	LinkID    uint      `gorm:"not null;uniqueIndex:idx_clickstat_unique"`
	Country   string    `gorm:"uniqueIndex:idx_clickstat_unique"`
	Device    string    `gorm:"uniqueIndex:idx_clickstat_unique"`
	Browser   string    `gorm:"uniqueIndex:idx_clickstat_unique"`
	Referrer  string    `gorm:"uniqueIndex:idx_clickstat_unique"`
	Day       time.Time `gorm:"not null;uniqueIndex:idx_clickstat_unique"`
	Count     uint      `gorm:"not null;default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
