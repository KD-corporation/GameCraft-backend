package helpers

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string {
	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	// Generate a 6-digit number
	otp := rand.Intn(900000) + 100000 // ensures 100000–999999

	return strconv.Itoa(otp)
}
