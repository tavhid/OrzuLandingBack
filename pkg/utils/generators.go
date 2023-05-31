package utils

import (
	random "math/rand"
	"strconv"
	"time"
)

func GenOtp() (otp string) {
	random.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 4; i++ {
		otp += strconv.Itoa(random.Intn(10))
	}
	return
}
