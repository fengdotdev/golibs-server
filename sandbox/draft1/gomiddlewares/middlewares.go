package gomiddlewares

import (
	"context"
	"net/http"
)

const (
	MiddlewareKey = "GoMiddlewareKey"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type GoMiddleware struct {
	Name    string
	Handler HandleFunc
}

func NewGoMiddleware(name string, handler HandleFunc) *GoMiddleware {
	return &GoMiddleware{
		Name:    name,
		Handler: handler,
	}
}

func (gm *GoMiddleware) Middleware(next HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), MiddlewareKey, gm.Name)
		r = r.WithContext(ctx)
		gm.Handler(w, r)
		next(w, r)
	}
}
