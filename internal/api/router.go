package api

import "github.com/fasthttp/router"

type Router struct {
	rtr *router.Router
}

func NewRouter() *Router {
	r := &Router{
		rtr: router.New(),
	}
	return r
}
