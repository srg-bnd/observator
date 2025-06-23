package helpers

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func Sha256Hash(data string, key string) string {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(data))

	return hex.EncodeToString(hmac.Sum(nil))
}

func SecureRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
