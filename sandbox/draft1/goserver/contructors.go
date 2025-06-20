package goserver

import (
	"net/http"
	"time"
)

func NewGoServer() *GoServer {
	return &GoServer{
		handlers:     make(map[string]http.HandlerFunc),
		mux:          http.NewServeMux(),
		server:       nil,
		port:         8080,              // Default port
		readTimeout:  10 * time.Second,  // Default read timeout
		writeTimeout: 10 * time.Second,  // Default write timeout
		idleTimeout:  120 * time.Second, // Default idle timeout
	}
}
