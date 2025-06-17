// Verify Checksum
package middleware

import "net/http"

type Checksum struct {
	EncryptionKey string
}

func NewChecksum(encryptionKey string) *Checksum {
	return &Checksum{
		EncryptionKey: encryptionKey,
	}
}

func (c *Checksum) WithVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO
		next.ServeHTTP(w, r)
	})
}
