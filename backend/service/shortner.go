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

	for i := 0; i < maxRetries; i++ {
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
	for {
		b := make([]byte, 8)

		_, err := rand.Read(b)
		if err != nil {
			return "", err
		}

		// convert random bytes to numbers
		randNum := binary.BigEndian.Uint64(b)
		encoded := util.Base62Encode(int(randNum))

		// fix for annoying empty encoded value
		if encoded == "" {
			continue
		}

		if len(encoded) > length {
			encoded = encoded[:length]
		}

		return encoded, nil
	}
}
