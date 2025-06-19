// Verify Checksum
package middleware

import (
	"net/http"
)

type ChecksumBehaviour interface {
	Verify(string, string) error
}

type Checksum struct {
	ChecksumService ChecksumBehaviour
}

const ErrReadBody = "couldn't read the body during verify"

func NewChecksum(checksumService ChecksumBehaviour) *Checksum {
	return &Checksum{
		ChecksumService: checksumService,
	}
}

func (c *Checksum) WithVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	logger.Log.Info(ErrReadBody, zap.Error(err))
		// 	return
		// }

		// if r.Header.Get("HashSHA256") != "" && len(body) > 0 {
		// 	if err := c.ChecksumService.Verify(r.Header.Get("HashSHA256"), string(body)); err != nil {
		// 		logger.Log.Info("checksum verify:", zap.Error(err))

		// 		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		// 		return
		// 	}
		// }

		next.ServeHTTP(w, r)
	})
}
