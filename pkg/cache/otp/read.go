package otp

import (
	"cl/pkg/bootstrap/http/misc/response"
	"encoding/json"

	"github.com/go-redis/redis"
)

func (p *provider) Get(key string) (expandTo map[string]interface{}, err error) {
	var value string

	if err := p.rdb.Get(key).Scan(&value); err != nil {
		if err.Error() == redis.Nil.Error() {
			return map[string]interface{}{}, response.ErrDataNotFound
		}

		p.logger.Error(err)
		return nil, err
	}

	if err := json.Unmarshal([]byte(value), &expandTo); err != nil {
		p.logger.Error(err)
		return nil, err
	}

	return
}
