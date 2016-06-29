package main

import "net/http"

// MySecureHandler defines allowed hosts
type MySecureHandler struct {
	errorHandler http.Handler
}

// START OMIT

const (
	contentTypeHeader = "Content-Type:text/html"
	contentTypeValue  = "charset=UTF-8"
)

// ContentTypeMiddleware adds the content type to the response
func (s *MySecureHandler) ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add Content-Type header
		w.Header().Add(contentTypeHeader, contentTypeValue)

		// Respond to request
		next.ServeHTTP(w, r)
	})
}

// END OMIT
