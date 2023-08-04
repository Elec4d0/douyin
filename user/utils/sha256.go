package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(p string) string {
	// p: password
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(p))
	bytes := sha256Hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
