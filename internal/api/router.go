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
		port:    cfg.Server.Port,
		router:  rtr,
	}
	srv.Handler = rtr.Handler
	rtr.POST("/task", r.addTask)
	rtr.PUT("/task/image", r.addImage)

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

func (r *Router) addImage(ctx *fasthttp.RequestCtx) {
	data, err := ctx.MultipartForm()
	if err != nil {
		r.logger.Println(err)
		ctx.SetStatusCode(500)
		return
	}
	if len(data.Value["uuid"]) != 1 || len(data.File["image"]) != 1 {
		r.logger.Println("wrong input")
		ctx.SetStatusCode(500)
		return
	}
	img, err := data.File["image"][0].Open()
	if err != nil {
		r.logger.Println(err)
		ctx.SetStatusCode(500)
		return
	}

	imgBytes := make([]byte, data.File["image"][0].Size)
	uuid := data.Value["uuid"][0]
	_, err = img.Read(imgBytes)
	if err != nil {
		r.logger.Println(err)
		ctx.SetStatusCode(500)
		return
	}
	title := data.File["image"][0].Filename
	err = r.service.AddImageToTask(uuid, title, imgBytes)
	if err != nil {
		r.logger.Println(err)
		ctx.SetStatusCode(500)
		return
	}
	// err = r.service.AddImageToTask(data.Value["uuid"][0], imgBytes)
	// if err != nil {
	// 	ctx.SetStatusCode(500)
	// 	return
	// }
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
