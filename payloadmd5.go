package avlogparser

import (
	"crypto/md5"
	b64 "encoding/base64"
	"encoding/hex"
	"io"
)

// HashPayload : take in a string, properly hash it as a byte array
// and return the hex as string.
func HashPayload(payload string) string {
	h := md5.New()
	io.WriteString(h, payload)
	return hex.EncodeToString(h.Sum(nil))
}

// DecodeToString : return the decoded payload
func DecodeToString(payload string) string {
	sDec, _ := b64.StdEncoding.DecodeString(payload)
	return string(sDec)
}
