package main

import "net/http"

// MySecureHandler defines a secure handler
type MySecureHandler struct{}

// START OMIT

const (
	xFrameOptionsHeader     = "X-Frame-Options"
	xFrameOptionsDeny       = "DENY"
	xFrameOptionsSameOrigin = "SAMEORIGIN"
	xFrameOptionsFromURI    = "ALLOW-FROM http://www.example.com"
)

// XFrameOptionsMiddleware adds the X-Frame-Options header to the response
func (a *MySecureHandler) XFrameOptionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent page from being displayed in an iframe
		w.Header().Add(xFrameOptionsHeader, xFrameOptionsDeny)
		// OR
		// Prevent page from being displayed in an iframe on other sites
		w.Header().Add(xFrameOptionsHeader, xFrameOptionsSameOrigin)
		// OR
		// Allow page at specified URI to display page in an iframe
		// NOTE: Limited support in modern browsers
		w.Header().Add(xFrameOptionsHeader, xFrameOptionsFromURI)

		// Respond to request
		next.ServeHTTP(w, r)
	})
}

// END OMIT
