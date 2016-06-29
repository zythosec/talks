package main

import "net/http"

// MySecureHandler defines allowed hosts
type MySecureHandler struct{}

// START OMIT

const (
	contentTypeOptionsHeader = "X-Content-Type-Options"
	contentTypeOptionsValue  = "nosniff"
)

// ContentTypeOptionsMiddleware adds content type options headers
func (s *MySecureHandler) ContentTypeOptionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add X-Content-Type-Options header
		w.Header().Add(contentTypeOptionsHeader, contentTypeOptionsValue)

		// Respond to request
		next.ServeHTTP(w, r)
	})
}

// END OMIT
