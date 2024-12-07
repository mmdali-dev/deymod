package internal

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Encode(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
