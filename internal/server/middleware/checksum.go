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

// TODO: move checksumResponseWriter to personal package
type checksumResponseWriter struct {
	http.ResponseWriter
	setChecksumFunc func(*checksumResponseWriter, []byte)
	statusCode      int
	buf             *bytes.Buffer
}

func (crw *checksumResponseWriter) Write(b []byte) (int, error) {
	return crw.buf.Write(b)
}

func (crw *checksumResponseWriter) HackWrite() {
	crw.setChecksumFunc(crw, crw.buf.Bytes())

	crw.ResponseWriter.WriteHeader(crw.statusCode)
	crw.ResponseWriter.Write(crw.buf.Bytes())
}

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

		if expectedHash := r.Header.Get(hashKey); expectedHash != "" && len(body) > 0 {
			if err := c.ChecksumService.Verify(expectedHash, string(body)); err != nil {
				logger.Log.Info(ErrChecksumVerify.Error(), zap.String("hashKey", expectedHash), zap.Error(err))
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}

		cw := checksumResponseWriter{
			ResponseWriter: w,
			buf:            &bytes.Buffer{},
			setChecksumFunc: func(crw *checksumResponseWriter, b []byte) {
				if len(b) == 0 {
					return
				}

				checksum, err := c.ChecksumService.Sum(string(b))
				if err != nil {
					logger.Log.Info(ErrChecksumVerify.Error(), zap.Error(err))
					return
				}

				crw.Header().Set(hashKey, checksum)
			},
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(&cw, r)
		cw.HackWrite()
	})
}
