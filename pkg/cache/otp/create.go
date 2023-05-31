package otp

import (
	"cl/pkg/bootstrap/http/misc/response"
	"encoding/json"
	"time"
)

func (p *provider) Save(key string, value interface{}) (err error) {
	// Convert payload data to string
	payload, err := json.Marshal(value)
	if err != nil {
		p.logger.Error("On parsing", err.Error())
		return response.ErrSomethingWentWrong
	}

	// Save data
	status := p.rdb.Set(key, string(payload), time.Minute*time.Duration(30))
	if status.Err() != nil {
		p.logger.Error("Error on save data", status.Err())
		return response.ErrSomethingWentWrong
	}

	return
}
