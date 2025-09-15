package helpers

import (
	"math/rand"
	"time"
)

func OptGenerate() string {
	s := "0123456789"
	rand.Seed(time.Now().UnixNano()) // randomize seed
	result := make([]byte, 6)
	for i := 0; i < 6; i++ {
		result[i] = s[rand.Intn(len(s))]
	}
	return string(result)
}