package service

import (
	"fmt"
	"log"
	"net/netip"
	"sync"
	"sync/atomic"
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
	DB             *gorm.DB
	GeoIPDB        *geoip2.Reader
	jobs           chan ClickJob
	done           chan struct{}
	wg             sync.WaitGroup
	batchSize      int
	flushInterval  time.Duration
	processedCount uint64
	errorCount     uint64
	droppedCount   uint64
}

func NewClickService(db *gorm.DB, geoipPath string, numWorkers int) *ClickService {
	geoIPDB, err := geoip2.Open(geoipPath)
	if err != nil {
		log.Fatalf("failed to open GeoIP database: %v", err)
	}

	s := &ClickService{
		DB:            db,
		GeoIPDB:       geoIPDB,
		jobs:          make(chan ClickJob, 1000),
		done:          make(chan struct{}),
		batchSize:     100,
		flushInterval: 5 * time.Second,
	}

	// Start worker pool
	for i := 0; i < numWorkers; i++ {
		s.wg.Add(1)
		go s.batchWorker(i)
	}

	log.Printf("ClickService started with %d workers", numWorkers)
	return s
}

func (s *ClickService) batchWorker(id int) {
	defer s.wg.Done()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Click worker %d panicked: %v", id, r)
		}
	}()

	log.Printf("Click worker %d started", id)
	batch := make([]ClickJob, 0, s.batchSize)
	ticker := time.NewTicker(s.flushInterval)
	defer ticker.Stop()

	for {
		select {
		case job, ok := <-s.jobs:
			if !ok {
				// Channel closed, process remaining batch and exit
				if len(batch) > 0 {
					s.processBatch(batch)
				}
				log.Printf("Click worker %d stopped", id)
				return
			}

			batch = append(batch, job)
			if len(batch) >= s.batchSize {
				s.processBatch(batch)
				batch = batch[:0]
			}

		case <-ticker.C:
			if len(batch) > 0 {
				s.processBatch(batch)
				batch = batch[:0]
			}

		case <-s.done:
			// Process remaining clicks before stopping
			if len(batch) > 0 {
				s.processBatch(batch)
			}
			log.Printf("Click worker %d stopped", id)
			return
		}
	}
}

func (s *ClickService) processBatch(jobs []ClickJob) {
	if len(jobs) == 0 {
		return
	}

	// Process entire batch in a single transaction for better performance
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		for _, job := range jobs {
			if err := s.processClickInTx(tx, job); err != nil {
				// Log but continue processing other clicks in batch
				log.Printf("Error processing click for link %d: %v", job.LinkID, err)
				atomic.AddUint64(&s.errorCount, 1)
			} else {
				atomic.AddUint64(&s.processedCount, 1)
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("Batch transaction error: %v", err)
	}
}

func (s *ClickService) processClickInTx(tx *gorm.DB, job ClickJob) error {
	// Parse User-Agent
	ua := useragent.New(job.UserAgentStr)
	browserName, _ := ua.Browser()
	device := "Desktop"
	if ua.Mobile() {
		device = "Mobile"
	} else if ua.Bot() {
		device = "Bot"
	}

	// Filter out bot traffic
	if device == "Bot" {
		log.Printf("Filtered bot click for link %d", job.LinkID)
		return nil
	}

	// Normalize country
	country := "Unknown"
	if c, err := s.LookupCountry(job.IP); err == nil && c != "" {
		country = c
	}

	// Normalize referrer
	referrer := job.Referrer
	if referrer == "" {
		referrer = "direct"
	}

	day := time.Now().UTC().Truncate(24 * time.Hour)
	now := time.Now()

	// Store raw click event
	click := model.Click{
		LinkID:   job.LinkID,
		Country:  country,
		Device:   device,
		Browser:  browserName,
		Referrer: referrer,
		IP:       job.IP,
	}
	if err := tx.Create(&click).Error; err != nil {
		return fmt.Errorf("raw click: %w", err)
	}

	// Update daily stats
	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "link_id"}, {Name: "day"}},
		DoUpdates: clause.Assignments(map[string]any{
			"clicks":     gorm.Expr("daily_link_stats.clicks + 1"),
			"updated_at": now,
		}),
	}).Create(&model.DailyLinkStats{
		LinkID: job.LinkID,
		Day:    day,
		Clicks: 1,
	}).Error; err != nil {
		return fmt.Errorf("daily stats: %w", err)
	}

	// Update country stats
	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "link_id"}, {Name: "country"}},
		DoUpdates: clause.Assignments(map[string]any{
			"clicks":     gorm.Expr("link_country_stats.clicks + 1"),
			"updated_at": now,
		}),
	}).Create(&model.LinkCountryStats{
		LinkID:  job.LinkID,
		Country: country,
		Clicks:  1,
	}).Error; err != nil {
		return fmt.Errorf("country stats: %w", err)
	}

	// Update device stats
	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "link_id"}, {Name: "device"}},
		DoUpdates: clause.Assignments(map[string]any{
			"clicks":     gorm.Expr("link_device_stats.clicks + 1"),
			"updated_at": now,
		}),
	}).Create(&model.LinkDeviceStats{
		LinkID: job.LinkID,
		Device: device,
		Clicks: 1,
	}).Error; err != nil {
		return fmt.Errorf("device stats: %w", err)
	}

	// Update browser stats
	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "link_id"}, {Name: "browser"}},
		DoUpdates: clause.Assignments(map[string]any{
			"clicks":     gorm.Expr("link_browser_stats.clicks + 1"),
			"updated_at": now,
		}),
	}).Create(&model.LinkBrowserStats{
		LinkID:  job.LinkID,
		Browser: browserName,
		Clicks:  1,
	}).Error; err != nil {
		return fmt.Errorf("browser stats: %w", err)
	}

	// Update referrer stats
	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "link_id"}, {Name: "referrer"}},
		DoUpdates: clause.Assignments(map[string]any{
			"clicks":     gorm.Expr("link_referrer_stats.clicks + 1"),
			"updated_at": now,
		}),
	}).Create(&model.LinkReferrerStats{
		LinkID:   job.LinkID,
		Referrer: referrer,
		Clicks:   1,
	}).Error; err != nil {
		return fmt.Errorf("referrer stats: %w", err)
	}

	return nil
}

func (s *ClickService) GetStats() (processed, errors, dropped uint64) {
	return atomic.LoadUint64(&s.processedCount),
		atomic.LoadUint64(&s.errorCount),
		atomic.LoadUint64(&s.droppedCount)
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
		atomic.AddUint64(&s.droppedCount, 1)
		log.Printf("Click job queue full, dropping click for link %d (total dropped: %d)",
			linkID, atomic.LoadUint64(&s.droppedCount))
	}
}

func (s *ClickService) Close() {
	log.Println("Shutting down ClickService...")

	// Signal workers to finish current work
	close(s.done)

	// Stop accepting new jobs
	close(s.jobs)

	// Wait for all workers to complete with timeout
	waitCh := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(waitCh)
	}()

	select {
	case <-waitCh:
		processed, errors, dropped := s.GetStats()
		log.Printf("ClickService shutdown complete. Processed: %d, Errors: %d, Dropped: %d",
			processed, errors, dropped)
	case <-time.After(30 * time.Second):
		processed, errors, dropped := s.GetStats()
		log.Printf("ClickService shutdown timeout. Processed: %d, Errors: %d, Dropped: %d",
			processed, errors, dropped)
	}

	// Close GeoIP database
	if s.GeoIPDB != nil {
		if err := s.GeoIPDB.Close(); err != nil {
			log.Printf("Error closing GeoIP database: %v", err)
		}
	}

	log.Println("ClickService stopped")
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
