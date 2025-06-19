package gogroup

import (
	"net/http"

	"github.com/fengdotdev/golibs-server/sandbox/draft1/gomiddlewares"
)

type Group struct {
	Name        string
	Handlers    []http.HandlerFunc
	Middlewares []func(http.Handler) http.Handler
	Methods     []string
}

func NewGoGroup(name string) *Group {
	return &Group{
		Name:        name,
		Handlers:    make([]http.HandlerFunc, 0),
		Middlewares: make([]func(http.Handler) http.Handler, 0),
		Methods:     make([]string, 0),
	}
}

func (g *Group) AddHandler(method string, handler http.HandlerFunc) {
	g.Methods = append(g.Methods, method)
	g.Handlers = append(g.Handlers, handler)
}

func (g *Group) AddMiddleware(middleware func(http.Handler) http.Handler) {
	g.Middlewares = append(g.Middlewares, middleware)
}
func (g *Group) AddGoMiddleware(middleware gomiddlewares.GoMiddleware) {
	g.Middlewares = append(g.Middlewares, middleware.Middleware(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			middleware.Handler(w, r)
			next.ServeHTTP(w, r)
		})
	}))
}
