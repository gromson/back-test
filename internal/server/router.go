package server

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type middleware func(handle httprouter.Handle) httprouter.Handle

type router struct {
	*httprouter.Router
}

func newRouter() *router {
	return &router{
		httprouter.New(),
	}
}

func (r *router) newGroup() *routeGroup {
	return &routeGroup{
		router:      r,
		middlewares: make([]middleware, 0, 2),
	}
}

type routeGroup struct {
	router      *router
	middlewares []middleware
}

func (g *routeGroup) GET(path string, handle httprouter.Handle) {
	g.router.GET(path, g.wrap(handle))
}

func (g *routeGroup) POST(path string, handle httprouter.Handle) {
	g.router.POST(path, g.wrap(handle))
}

func (g *routeGroup) PUT(path string, handle httprouter.Handle) {
	g.router.PUT(path, g.wrap(handle))
}

func (g *routeGroup) AddMiddleware(m middleware) {
	g.middlewares = append(g.middlewares, m)
}

func (g *routeGroup) wrap(handler httprouter.Handle) httprouter.Handle {
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