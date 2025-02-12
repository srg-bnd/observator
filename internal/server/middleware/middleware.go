// Middleware for server
package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func CheckMethodPost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resWriter http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(resWriter, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(resWriter, req)
	})
}
