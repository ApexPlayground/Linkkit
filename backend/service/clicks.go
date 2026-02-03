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
	"gorm.io/gorm/clause"
)

type ClickJob struct {
	LinkID       uint
	IP           string
	UserAgentStr string
	Referrer     string
}

type ClickService struct {
	DB      *gorm.DB
	GeoIPDB *geoip2.Reader
	jobs    chan ClickJob
	done    chan struct{}
}

func NewClickService(db *gorm.DB, geoipPath string, numWorkers int) *ClickService {
	geoIPDB, err := geoip2.Open(geoipPath)
	if err != nil {
		log.Fatalf("failed to open GeoIP database: %v", err)
	}

	s := &ClickService{
		DB:      db,
		GeoIPDB: geoIPDB,
		jobs:    make(chan ClickJob, 1000), // Buffer for 1000 pending clicks
		done:    make(chan struct{}),
	}

	// Start worker pool
	for i := 0; i < numWorkers; i++ {
		go s.worker(i)
	}

	return s
}

func (s *ClickService) worker(id int) {
	log.Printf("Click worker %d started", id)
	for {
		select {
		case job := <-s.jobs:
			s.processClick(job)
		case <-s.done:
			log.Printf("Click worker %d stopped", id)
			return
		}
	}
}

func (s *ClickService) processClick(job ClickJob) {
	// Parse User-Agent
	ua := useragent.New(job.UserAgentStr)
	browserName, _ := ua.Browser()

	device := "Desktop"
	if ua.Mobile() {
		device = "Mobile"
	} else if ua.Bot() {
		device = "Bot"
	}

	country := "Unknown"
	if c, err := s.LookupCountry(job.IP); err == nil {
		country = c
	}

	// Daily bucket (UTC)
	day := time.Now().UTC().Truncate(24 * time.Hour)

	stat := model.ClickStat{
		LinkID:   job.LinkID,
		Country:  country,
		Device:   device,
		Browser:  browserName,
		Referrer: job.Referrer,
		Day:      day,
		Count:    1,
	}

	err := s.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "link_id"},
			{Name: "country"},
			{Name: "device"},
			{Name: "browser"},
			{Name: "referrer"},
			{Name: "day"},
		},
		DoUpdates: clause.Assignments(map[string]any{
			"count":      gorm.Expr("click_stats.count + 1"),
			"updated_at": time.Now(),
		}),
	}).Create(&stat).Error

	if err != nil {
		log.Printf("Click tracking error for link %d: %v", job.LinkID, err)
	}
}

func (s *ClickService) TrackClick(linkID uint, ip, userAgentStr, referrer string) {
	job := ClickJob{
		LinkID:       linkID,
		IP:           ip,
		UserAgentStr: userAgentStr,
		Referrer:     referrer,
	}

	// Non-blocking send to job queue
	select {
	case s.jobs <- job:
		// Job queued successfully
	default:
		log.Printf("Click job queue full, dropping click for link %d", linkID)
	}
}

func (s *ClickService) Close() {
	// Signal all workers to stop
	close(s.done)

	// Wait a bit for workers to finish current jobs
	time.Sleep(2 * time.Second)

	// Close job channel
	close(s.jobs)

	// Close GeoIP database
	if s.GeoIPDB != nil {
		s.GeoIPDB.Close()
	}
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
