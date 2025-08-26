// Decryptor
package middleware

import "net/http"

type DecryptMethod interface {
	Decrypt([]byte) ([]byte, error)
}

type Decryptor struct {
	decryptMethod DecryptMethod
}

func NewDecryptor(decryptMethod DecryptMethod) *Decryptor {
	return &Decryptor{
		decryptMethod: decryptMethod,
	}
}

func (c *Decryptor) WithDecrypt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
