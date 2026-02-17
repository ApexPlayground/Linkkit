package model

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null;index"`
	LongUrl   string         `gorm:"not null"`
	ShortCode string         `gorm:"uniqueIndex"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID"`
}

type QRCode struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null;index"`
	OriginalURL string `gorm:"not null;size:2048"`
	QRImage     []byte `gorm:"type:bytea"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID"`
}

type Click struct {
	ID        uint      `gorm:"primaryKey"`
	LinkID    uint      `gorm:"not null;index:idx_link_created"`
	QRID      uint      `gorm:"not null;index:idx_link_created"`
	Country   string    `gorm:"size:50"`
	Device    string    `gorm:"size:20"`
	Browser   string    `gorm:"size:50"`
	Referrer  string    `gorm:"size:500"`
	IP        string    `gorm:"size:45;index"`
	CreatedAt time.Time `gorm:"not null;index:idx_link_created"`
}

type DailyLinkStats struct {
	ID        uint      `gorm:"primaryKey"`
	LinkID    uint      `gorm:"not null;uniqueIndex:idx_daily_link"`
	Day       time.Time `gorm:"not null;type:date;uniqueIndex:idx_daily_link"`
	Clicks    uint      `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type LinkCountryStats struct {
	ID        uint   `gorm:"primaryKey"`
	LinkID    uint   `gorm:"not null;uniqueIndex:idx_link_country"`
	Country   string `gorm:"not null;size:50;uniqueIndex:idx_link_country"`
	Clicks    uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type LinkDeviceStats struct {
	ID        uint   `gorm:"primaryKey"`
	LinkID    uint   `gorm:"not null;uniqueIndex:idx_link_device"`
	Device    string `gorm:"not null;size:20;uniqueIndex:idx_link_device"`
	Clicks    uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type LinkBrowserStats struct {
	ID        uint   `gorm:"primaryKey"`
	LinkID    uint   `gorm:"not null;uniqueIndex:idx_link_browser"`
	Browser   string `gorm:"not null;size:50;uniqueIndex:idx_link_browser"`
	Clicks    uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type LinkReferrerStats struct {
	ID        uint   `gorm:"primaryKey"`
	LinkID    uint   `gorm:"not null;uniqueIndex:idx_link_referrer"`
	Referrer  string `gorm:"not null;size:500;uniqueIndex:idx_link_referrer"`
	Clicks    uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type DailyQRStats struct {
	ID        uint      `gorm:"primaryKey"`
	QRID      uint      `gorm:"not null;uniqueIndex:idx_daily_qr"`
	Day       time.Time `gorm:"not null;type:date;uniqueIndex:idx_daily_qr"`
	Scans     uint      `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type QRCountryStats struct {
	ID        uint   `gorm:"primaryKey"`
	QRID      uint   `gorm:"not null;uniqueIndex:idx_qr_country"`
	Country   string `gorm:"not null;size:50;uniqueIndex:idx_qr_country"`
	Scans     uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type QRDeviceStats struct {
	ID        uint   `gorm:"primaryKey"`
	QRID      uint   `gorm:"not null;uniqueIndex:idx_qr_device"`
	Device    string `gorm:"not null;size:20;uniqueIndex:idx_qr_device"`
	Scans     uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type QRBrowserStats struct {
	ID        uint   `gorm:"primaryKey"`
	QRID      uint   `gorm:"not null;uniqueIndex:idx_qr_browser"`
	Browser   string `gorm:"not null;size:50;uniqueIndex:idx_qr_browser"`
	Scans     uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type QRReferrerStats struct {
	ID        uint   `gorm:"primaryKey"`
	QRID      uint   `gorm:"not null;uniqueIndex:idx_qr_referrer"`
	Referrer  string `gorm:"not null;size:500;uniqueIndex:idx_qr_referrer"`
	Scans     uint   `gorm:"not null;default:0"`
	UpdatedAt time.Time
}
