package common

import (
	"crypto/sha256"
	"fmt"
)

func Hash(message string) string {
	byte := []byte(message)
	hash := fmt.Sprintf("%x", sha256.Sum256(byte))
	return hash
}
