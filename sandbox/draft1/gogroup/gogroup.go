package gogroup

import (
	"net/http"

	"github.com/fengdotdev/golibs-server/sandbox/draft1/gomiddlewares"
)

type Group struct {
	Name        string
	Handlers    []http.HandlerFunc
	Middlewares []func(http.HandlerFunc) http.HandlerFunc
	Methods     []string
}

func NewGoGroup(name string) *Group {
	return &Group{
		Name:        name,
		Handlers:    make([]http.HandlerFunc, 0),
		Middlewares: make([]func(http.HandlerFunc) http.HandlerFunc, 0),
	}
}

func (g *Group) AddHandler(method string, handler http.HandlerFunc) {
	g.Methods = append(g.Methods, method)
	g.Handlers = append(g.Handlers, handler)
}

func (g *Group) AddMiddleware(middleware func(http.HandlerFunc) http.HandlerFunc) {
	g.Middlewares = append(g.Middlewares, middleware)
}
func (g *Group) AddGoMiddleware(middleware gomiddlewares.GoMiddleware) {

	middlewareFunc := middleware.Get()
	g.Middlewares = append(g.Middlewares, middlewareFunc)
}
