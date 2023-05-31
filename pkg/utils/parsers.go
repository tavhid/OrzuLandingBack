package utils

import (
	"cl/pkg/bootstrap/http/misc/response"
	"time"
)

// ParseOTP ...
func ParseOTP(payload map[string]interface{}) (map[string]interface{}, error) {
	if len(payload) == 0 {
		payload = make(map[string]interface{})
		payload["created_at"] = time.Now().Unix()
		payload["attempts"] = 1
	} else {
		createdAt, _ := payload["created_at"].(float64)
		payload["created_at"] = int64(createdAt)
		attempts, _ := payload["attempts"].(float64)
		payload["attempts"] = int(attempts)
	}

	if (payload["attempts"]).(int) > 3 {
		return nil, response.ErrLimitExceeded
	}

	return payload, nil
}
