package otp

func (p *provider) Delete(key string) {
	p.rdb.Del(key)
}