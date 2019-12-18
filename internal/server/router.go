package server

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Middleware func(handle httprouter.Handle) httprouter.Handle

type Router struct {
	*httprouter.Router
}

func NewRouter() *Router {
	return &Router{
		httprouter.New(),
	}
}

func (r *Router) NewGroup() *RouteGroup {
	return &RouteGroup{
		router:      r,
		middlewares: make([]Middleware, 0, 2),
	}
}

type RouteGroup struct {
	router      *Router
	middlewares []Middleware
}

func (g *RouteGroup) GET(path string, handle httprouter.Handle) {
	g.router.GET(path, g.wrap(handle))
}

func (g *RouteGroup) POST(path string, handle httprouter.Handle) {
	g.router.POST(path, g.wrap(handle))
}

func (g *RouteGroup) PUT(path string, handle httprouter.Handle) {
	g.router.PUT(path, g.wrap(handle))
}

func (g *RouteGroup) AddMiddleware(m Middleware) {
	g.middlewares = append(g.middlewares, m)
}

func (g *RouteGroup) wrap(handler httprouter.Handle) httprouter.Handle {
	h := handler

	for _, m := range g.middlewares {
		h = m(h)
	}

	return h
}

func panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Printf("[%s %s] %v", r.Method, r.URL.Path, err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}