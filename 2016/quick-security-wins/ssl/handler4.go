package main

import "net/http"

// START OMIT

func (a *MySecureHandler) HTTPSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Determine if we are on HTTPS
		// 2. Do an SSL Check
		// 3. Strict Transport Security header. Only add header when we know it's an SSL connection.
		// 4. HPKP header.
		if len(s.opt.PublicKey) > 0 && isSSL {
			w.Header().Add(hpkpHeader, s.opt.PublicKey)
		}
		if err != nil {
			return
		}

		// Respond to request
		next.ServeHTTP(w, r)
	})
}

// END OMIT
