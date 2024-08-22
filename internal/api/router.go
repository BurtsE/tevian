package api

import (
	"bytes"
	"encoding/base64"
	"log"
	"tevian/internal/config"
	"tevian/internal/service"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Router struct {
	logger   *logrus.Logger
	service  service.Service
	router   *router.Router
	srv      *fasthttp.Server
	port     string
	user     string
	password string
}

func NewRouter(cfg *config.Config, service service.Service, logger *logrus.Logger) *Router {
	srv := fasthttp.Server{}
	rtr := router.New()
	r := &Router{
		logger:   logger,
		service:  service,
		srv:      &srv,
		port:     cfg.Server.Port,
		router:   rtr,
		user:     cfg.Credentials.Login,
		password: cfg.Credentials.Password,
	}
	srv.Handler = r.basicAuth(rtr.Handler)

	rtr.GET("/task/{uuid}", r.task)
	rtr.POST("/task", r.addTask)
	rtr.POST("/task/{uuid}/start", r.startTask)
	rtr.DELETE("/task/{uuid}", r.deleteTask)
	rtr.PUT("/task/image", r.addImage)

	r.router.GET("/status", statusHandler)
	return r
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

func (r *Router) basicAuth(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		auth := ctx.Request.Header.Peek("Authorization")
		if !r.checkAuthReq(auth) {
			ctx.SetStatusCode(401)
			ctx.SetBody([]byte("authorization failed"))
			return
		}
		handler(ctx)
		r.logger.Printf("api request: %s ;status code: %d", ctx.Path(), ctx.Response.StatusCode())
	}
}

func (r *Router) checkAuthReq(auth []byte) bool {
	i := bytes.IndexByte(auth, ' ')
	if i == -1 {
		return false
	}

	if !bytes.EqualFold(auth[:i], []byte("basic")) {
		return false
	}
	log.Println(string(auth[i+1:]))
	decoded, err := base64.StdEncoding.DecodeString(string(auth[i+1:]))
	if err != nil {
		return false
	}

	credentials := bytes.Split(decoded, []byte(":"))
	if len(credentials) <= 1 {
		log.Println(string(credentials[0]))
		return false
	}

	user := string(credentials[0])
	pass := string(credentials[1])
	if user != r.user || pass != r.password {
		return false
	}
	return true
}
