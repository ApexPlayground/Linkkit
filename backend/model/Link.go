package model

import "time"

type Link struct {
	ID        uint `gorm:"primaryKey"`
	LongUrl   string
	ShortCode string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

type Click struct {
	ID        uint `gorm:"primaryKey"`
	LinkID    uint `gorm:"not null;index"`
	CreatedAt time.Time
	IP        string
	UserAgent string
	Referrer  string
}
