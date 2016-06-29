package main

import (
	"fmt"
	"net/http"
	"strings"
)

// MySecureHandler defines allowed hosts
type MySecureHandler struct {
	options      Options
	errorHandler http.Handler
}

func defaultErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error Serving Response", http.StatusInternalServerError)
}

// New creates a new MySecureHandler
func New(options Options) *MySecureHandler {
	return &MySecureHandler{
		options:      options,
		errorHandler: http.HandlerFunc(defaultErrorHandler),
	}
}

// START OMIT

// AllowedHostsMiddleware handles the list of allowed hosts
func (a *MySecureHandler) AllowedHostsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		if len(a.options.AllowedHosts) > 0 {
			isGoodHost := false
			for _, allowedHost := range a.options.AllowedHosts {
				if strings.EqualFold(allowedHost, r.Host) {
					isGoodHost = true
					break
				}
			}
			if !isGoodHost {
				a.errorHandler.ServeHTTP(w, r)
				err = fmt.Errorf("Bad host name: %s", r.Host)
			}
		}
		// If there was an error, do not continue request
		if err != nil {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// END OMIT
