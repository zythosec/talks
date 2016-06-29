package main

import (
	"fmt"
	"net/http"
)

// START OMIT

func (a *MySecureHandler) HTTPSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Determine if we are on HTTPS.
		// 2. SSL check.
		if a.options.SSLRedirect && !isSSL {
			url := r.URL
			url.Scheme = "https"
			url.Host = r.Host

			if len(a.options.SSLHost) > 0 {
				url.Host = a.options.SSLHost
			}

			status := http.StatusMovedPermanently
			if a.options.SSLTemporaryRedirect {
				status = http.StatusTemporaryRedirect
			}

			http.Redirect(w, r, url.String(), status)
			err = fmt.Errorf("Redirecting to HTTPS")
		}
	})
}

// END OMIT
