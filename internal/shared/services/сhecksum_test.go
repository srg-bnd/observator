package services

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	randomString, err := secureRandomString(16)
	assert.Nil(t, err)

	data := "test data"
	hmac := hmac.New(sha256.New, []byte(randomString))
	hmac.Write([]byte(data))

	err = NewChecksum(randomString).Verify(string(hmac.Sum(nil)), data)
	assert.Nil(t, err)
}

func TestSum(t *testing.T) {
	randomString, err := secureRandomString(16)
	assert.Nil(t, err)

	data := "test data"
	hmac := hmac.New(sha256.New, []byte(randomString))
	hmac.Write([]byte(data))
	sum, err := NewChecksum(randomString).Sum(data)
	assert.Nil(t, err)

	assert.Equal(t, string(hmac.Sum(nil)), sum)
}

// Helpers

func secureRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
