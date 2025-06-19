package gomiddlewares

import (
	"context"
	"net/http"
)

const (
	MiddlewareKey = "GoMiddlewareKey"
)

type GoMiddleware struct {
	Name    string
	Handler http.HandlerFunc
}

func NewGoMiddleware(name string, handler http.HandlerFunc) *GoMiddleware {
	return &GoMiddleware{
		Name:    name,
		Handler: handler,
	}
}

func (gm *GoMiddleware) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), MiddlewareKey, gm.Name)
		r = r.WithContext(ctx)
		gm.Handler(w, r)
		next(w, r)
	}
}

func (gm *GoMiddleware) Get() func(http.HandlerFunc) http.HandlerFunc {
	return gm.Middleware
}
