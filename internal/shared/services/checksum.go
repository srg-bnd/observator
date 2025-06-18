package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
)

var ErrVerify = errors.New("unverified data")

type Checksum struct {
	hmac hash.Hash
}

func NewChecksum(secretkey string) *Checksum {
	return &Checksum{
		hmac: hmac.New(sha256.New, []byte(secretkey)),
	}
}

func (c *Checksum) Verify(dataHash, data string) error {
	sum, err := c.Sum(data)
	if err != nil {
		return err
	}

	if sum != dataHash {
		return ErrVerify
	}

	return nil
}

func (c *Checksum) Sum(data string) (string, error) {
	defer c.hmac.Reset()

	_, err := c.hmac.Write([]byte(data))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(c.hmac.Sum(nil)), nil
}
