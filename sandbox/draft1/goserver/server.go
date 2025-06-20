package goserver

import (
	"net/http"
	"time"
)

type GoServer struct {
	handlers     map[string]http.HandlerFunc
	mux          *http.ServeMux
	server       *http.Server
	port         int
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
	secure       bool
}
