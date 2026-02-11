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

// Raw click events - stores individual clicks
type Click struct {
	ID        uint      `gorm:"primaryKey"`
	LinkID    uint      `gorm:"not null;index:idx_link_created"`
	Country   string    `gorm:"size:50"`
	Device    string    `gorm:"size:20"`
	Browser   string    `gorm:"size:50"`
	Referrer  string    `gorm:"size:500"`
	IP        string    `gorm:"size:45;index"`
	CreatedAt time.Time `gorm:"not null;index:idx_link_created"`
}

// Aggregated stats, one row per link per day
type DailyLinkStats struct {
	ID        uint      `gorm:"primaryKey"`
	LinkID    uint      `gorm:"not null;uniqueIndex:idx_daily_link"`
	Day       time.Time `gorm:"not null;type:date;uniqueIndex:idx_daily_link"`
	Clicks    uint      `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

// Country breakdown, one row per link per country
type LinkCountryStats struct {
	ID        uint   `gorm:"primaryKey"`
	LinkID    uint   `gorm:"not null;uniqueIndex:idx_link_country"`
	Country   string `gorm:"not null;size:50;uniqueIndex:idx_link_country"`
	Clicks    uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

// Device breakdown, one row per link per device type
type LinkDeviceStats struct {
	ID        uint   `gorm:"primaryKey"`
	LinkID    uint   `gorm:"not null;uniqueIndex:idx_link_device"`
	Device    string `gorm:"not null;size:20;uniqueIndex:idx_link_device"`
	Clicks    uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

// Browser breakdown, one row per link per browser
type LinkBrowserStats struct {
	ID        uint   `gorm:"primaryKey"`
	LinkID    uint   `gorm:"not null;uniqueIndex:idx_link_browser"`
	Browser   string `gorm:"not null;size:50;uniqueIndex:idx_link_browser"`
	Clicks    uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

// Referrer breakdown, one row per link per referrer
type LinkReferrerStats struct {
	ID        uint   `gorm:"primaryKey"`
	LinkID    uint   `gorm:"not null;uniqueIndex:idx_link_referrer"`
	Referrer  string `gorm:"not null;size:500;uniqueIndex:idx_link_referrer"`
	Clicks    uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}
