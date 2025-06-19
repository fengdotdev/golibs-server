package goserver

import "net/http"


type GoServer struct {
	handlers map[string]http.HandlerFunc
}
