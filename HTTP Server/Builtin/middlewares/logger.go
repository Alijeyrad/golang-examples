package middleware

import (
	"log"
	"net/http"
)

func Logger() Middleware {
	return func(next Handler) Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
			next(w, r)
		}
	}
}
