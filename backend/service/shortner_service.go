package service

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/ApexPlayground/Linkkit/config"
	"github.com/ApexPlayground/Linkkit/model"
	"github.com/ApexPlayground/Linkkit/util"
)

type ServiceError struct {
	Status  int
	Message string
}

func (e ServiceError) Error() string {
	return e.Message
}

func CreateShortLink(longUrl string) (model.Link, error) {
	const (
		codeLength = 7
		maxRetries = 5
	)

	for range maxRetries {
		code, err := GenerateShortCode(codeLength)
		if err != nil {
			return model.Link{}, ServiceError{
				Status:  500,
				Message: "failed to generate shortcode",
			}
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

		return model.Link{}, ServiceError{
			Status:  500,
			Message: fmt.Sprintf("could not save link: %v", err),
		}
	}
	return model.Link{}, ServiceError{
		Status:  500,
		Message: "could not generate unique shortcode after retries",
	}
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
		rand_num := binary.BigEndian.Uint64(b)
		encoded := util.Base62Encode(int(rand_num))

		if encoded == "" {
			continue
		}

		if len(encoded) > length {
			encoded = encoded[:length]
		}

		return encoded, nil
	}
}
