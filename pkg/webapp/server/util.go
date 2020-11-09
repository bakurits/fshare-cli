package server

import (
	"crypto/rand"
	"encoding/base64"
)

func randToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
