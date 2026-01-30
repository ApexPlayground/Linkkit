package service

import (
	"fmt"
	"log"
	"net/netip"
	"time"

	"github.com/ApexPlayground/Linkkit/model"
	"github.com/mssola/useragent"
	"github.com/oschwald/geoip2-golang/v2"
	"gorm.io/gorm"
)

type ClickService struct {
	DB      *gorm.DB
	GeoIPDB *geoip2.Reader
}

func NewClickService(db *gorm.DB, geoipPath string) *ClickService {
	geoIPDB, err := geoip2.Open(geoipPath)
	if err != nil {
		log.Fatalf("failed to open GeoIP database: %v", err)
	}

	return &ClickService{
		DB:      db,
		GeoIPDB: geoIPDB,
	}
}

func (s *ClickService) Close() {
	if s.GeoIPDB != nil {
		s.GeoIPDB.Close()
	}
}

func (s *ClickService) TrackClick(linkID uint, ip, userAgentStr, referrer string) {
	go func() {
		// Parse User-Agent string
		ua := useragent.New(userAgentStr)
		browserName, _ := ua.Browser()
		device := "Desktop"
		if ua.Mobile() {
			device = "Mobile"
		} else if ua.Bot() {
			device = "Bot"
		}

		// Lookup country
		country := "Unknown"
		if c, err := s.LookupCountry(ip); err == nil {
			country = c
		}

		click := model.Click{
			LinkID:    linkID,
			CreatedAt: time.Now(),
			IP:        ip,
			UserAgent: userAgentStr,
			Referrer:  referrer,
			Browser:   browserName,
			Device:    device,
			Country:   country,
		}

		if err := s.DB.Create(&click).Error; err != nil {
			log.Printf("Click tracking error for link %d: %v", linkID, err)
		}
	}()
}

func (s *ClickService) LookupCountry(ipStr string) (string, error) {
	ip, err := netip.ParseAddr(ipStr)
	if err != nil {
		return "", fmt.Errorf("invalid IP: %w", err)
	}

	record, err := s.GeoIPDB.Country(ip)
	if err != nil {
		return "", fmt.Errorf("GeoIP lookup failed: %w", err)
	}

	if !record.HasData() || record.Country.ISOCode == "" {
		return "", fmt.Errorf("no country data for IP %s", ipStr)
	}

	return record.Country.Names.English, nil
}
