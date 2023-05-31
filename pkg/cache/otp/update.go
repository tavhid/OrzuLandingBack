package otp

func (p *provider) Update(key string, value interface{}) (err error) {
	if _, err = p.Get(key); err != nil {
		return
	}

	if err = p.Save(key, value); err != nil {
		return
	}

	return
}