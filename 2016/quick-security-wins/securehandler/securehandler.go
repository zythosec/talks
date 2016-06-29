package main

import (
	"net/http"
)

// START OMIT

// Options defines our security options
type Options struct {
}

// MySecureHandler defines allowed hosts
type MySecureHandler struct {
	options      Options
	errorHandler http.Handler
}

func defaultErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error Serving Response", http.StatusInternalServerError)
}

// NewSecureHandler creates a new MySecureHandler
func NewSecureHandler(options Options) *MySecureHandler {
	return &MySecureHandler{
		options:      options,
		errorHandler: http.HandlerFunc(defaultErrorHandler),
	}
}

// END OMIT
