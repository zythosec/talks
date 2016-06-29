package main

import (
	"fmt"
	"net/http"
)

// START OMIT
func (a *MySecureHandler) HTTPSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Determine if we are on HTTPS
		// 2. Do an SSL Check
		// 3. Strict Transport Security header. Only add header when we know it's an SSL connection.
		if a.options.STSSeconds != 0 && (isSSL || a.options.ForceSTSHeader) {
			stsSub := ""
			if a.options.STSIncludeSubdomains {
				stsSub = stsSubdomainString
			}

			if a.options.STSPreload {
				stsSub += stsPreloadString
			}

			w.Header().Add(stsHeader, fmt.Sprintf("max-age=%d%s", a.options.STSSeconds, stsSub))
		}

		if err != nil {
			return
		}
	})
}

// END OMIT
