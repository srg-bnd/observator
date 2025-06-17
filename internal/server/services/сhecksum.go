package services

import (
	"crypto/hmac"
	"crypto/sha256"
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
	_, err := c.hmac.Write([]byte(data))
	if err != nil {
		return err
	}

	if string(c.hmac.Sum(nil)) != dataHash {
		return ErrVerify
	}

	return nil
}
