package api

import (
	"tevian/internal/config"
	"tevian/internal/service"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Router struct {
	logger  *logrus.Logger
	service service.Service
	router  *router.Router
	srv     *fasthttp.Server
	port    string
}

func NewRouter(cfg *config.Config, service service.Service, logger *logrus.Logger) *Router {
	srv := fasthttp.Server{}
	rtr := router.New()
	r := &Router{
		logger:  logger,
		service: service,
		srv:     &srv,
		port:    cfg.Postgres.Port,
		router:  rtr,
	}
	srv.Handler = rtr.Handler
	rtr.GET("/add-task", r.addTask)
	r.router.GET("/status", statusHandler)
	return r
}

func (r *Router) addTask(ctx *fasthttp.RequestCtx) {
	uuid, err := r.service.CreateTask()
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}
	ctx.Response.AppendBody([]byte(uuid))
}

func (r *Router) Start() error {
	return r.srv.ListenAndServe(r.port)
}
func (r *Router) Shutdown() error {
	return r.srv.Shutdown()
}

func statusHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
}
