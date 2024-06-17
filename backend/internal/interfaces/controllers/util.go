package controllers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func MakeRandomStr(length int) string {
	randBytes := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, randBytes)
	if err != nil {
		return ""
	}
	return base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
}

func Hash(char string, count int) string {
	hash := sha256.Sum256([]byte(char))
	for i := 1; i < count; i++ {
		hash = sha256.Sum256(hash[:])
	}
	return fmt.Sprintf("%x", hash)
}
