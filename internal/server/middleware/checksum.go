// Verify Checksum
package middleware

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/srg-bnd/observator/internal/server/logger"
	"go.uber.org/zap"
)

type ChecksumBehaviour interface {
	Verify(string, string) error
	Sum(string) (string, error)
}

type Checksum struct {
	ChecksumService ChecksumBehaviour
}

const hashKey = "HashSHA256"

var ErrChecksumVerify = errors.New("bad checksum verify")

func NewChecksum(checksumService ChecksumBehaviour) *Checksum {
	return &Checksum{
		ChecksumService: checksumService,
	}
}

func (c *Checksum) WithVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Log.Info(ErrChecksumVerify.Error(), zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r.Body.Close()

		if expectedChecksum := r.Header.Get(hashKey); expectedChecksum != "" && len(body) > 0 {
			if err := c.ChecksumService.Verify(expectedChecksum, string(body)); err != nil {
				logger.Log.Info(ErrChecksumVerify.Error(), zap.Error(err))
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)

		responseBody, err := io.ReadAll(r.Response.Body)
		if err != nil {
			logger.Log.Info(ErrChecksumVerify.Error(), zap.Error(err))
			return
		}

		checksum, err := c.ChecksumService.Sum(string(responseBody))
		if err != nil {
			logger.Log.Info(ErrChecksumVerify.Error(), zap.Error(err))
			return
		}

		w.Header().Set(hashKey, checksum)
	})
}
