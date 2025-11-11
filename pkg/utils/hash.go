package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5String returns the hex encoded MD5 hash of the provided string.
func MD5String(value string) string {
	sum := md5.Sum([]byte(value))
	return hex.EncodeToString(sum[:])
}

// SafeString trims whitespace and lowercases the string.
func SafeString(value string) string {
	return strings.TrimSpace(value)
}
