package service

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/model"
	"github.com/ApexPlayground/Linkkit/util"
)

func CreateShortLink(longUrl string) (model.Link, error) {
	const (
		codeLength = 7
		maxRetries = 5
		maxURLLen  = 2048
	)

	longUrl = strings.TrimSpace(longUrl)
	if len(longUrl) == 0 {
		return model.Link{}, fmt.Errorf("URL cannot be empty")
	}

	if len(longUrl) > maxURLLen {
		return model.Link{}, fmt.Errorf("URL too long")
	}

	parsed, err := url.ParseRequestURI(longUrl)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return model.Link{}, fmt.Errorf("invalid URL")
	}

	for range maxRetries {
		code, err := GenerateShortCode(codeLength)
		if err != nil {
			return model.Link{}, fmt.Errorf("failed to generate shortcode")
		}

		link := model.Link{
			LongUrl:   longUrl,
			ShortCode: code,
		}

		err = config.DB.Create(&link).Error
		if err == nil {
			return link, nil
		}

		// Retry ONLY on unique constraint violation
		if strings.Contains(err.Error(), "duplicate key") {
			continue
		}

		return model.Link{}, fmt.Errorf("could not save link: %v", err)
	}

	return model.Link{}, fmt.Errorf("could not generate unique shortcode after %d retries", maxRetries)
}

// GenerateShortCode generates a short code based on the long URL
func GenerateShortCode(length int) (string, error) {
	const maxAttempts = 10

	for range maxAttempts {
		b := make([]byte, 8)

		_, err := rand.Read(b)
		if err != nil {
			return "", fmt.Errorf("failed to generate random bytes: %v", err)
		}

		randNum := binary.BigEndian.Uint64(b)
		encoded := util.Base62Encode(int(randNum % (1 << 62)))

		// fix for annoying empty encoded value
		if encoded == "" {
			continue
		}

		if len(encoded) > length {
			encoded = encoded[:length]
		}

		return encoded, nil
	}
	return "", fmt.Errorf("failed to generate shortcode after %d attempts", maxAttempts)
}
