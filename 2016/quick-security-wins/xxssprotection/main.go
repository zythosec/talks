package main

import (
	"net/http"
)

// MySecureHandler defines allowed hosts
type MySecureHandler struct{}

// START OMIT

const (
	xssProtectionHeader = "X-XSS-Protection"
	xssProtectionValue  = "1; mode=block"
)

// XXssProtectionMiddleware adds X-XSS-Protection header to your responses
func (a *MySecureHandler) XXssProtectionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add X-XSS-Protection header
		w.Header().Add(xssProtectionHeader, xssProtectionValue)

		// Respond to request
		next.ServeHTTP(w, r)
	})
}

// END OMIT
