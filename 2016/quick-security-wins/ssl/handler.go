package main

import (
	"net/http"
	"strings"
)

// START OMIT

func (a *MySecureHandler) HTTPSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		// 1. Determine if we are on HTTPS.
		isSSL := strings.EqualFold(r.URL.Scheme, "https") || r.TLS != nil
		if !isSSL {
			for k, v := range a.options.SSLProxyHeaders {
				if r.Header.Get(k) == v {
					isSSL = true
					break
				}
			}
		}
	})
}

// END OMIT
