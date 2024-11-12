package api

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func (r *Router) task(ctx *fasthttp.RequestCtx) {
	var (
		uuid string
		ok   bool
	)
	if uuid, ok = ctx.UserValue("uuid").(string); !ok {
		r.logger.Println(ctx.UserValue("uuid"))
		ctx.SetStatusCode(500)
		return
	}
	task, err := r.service.Task(ctx, uuid)
	if err != nil {
		r.logger.Println(err)
		ctx.SetStatusCode(500)
		return
	}
	data, err := json.Marshal(&task)
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}
	ctx.Response.AppendBody(data)
}

func (r *Router) addTask(ctx *fasthttp.RequestCtx) {
	uuid, err := r.service.CreateTask(ctx)
	if err != nil {
		r.logger.Println(err)
		ctx.SetStatusCode(500)
		return
	}
	ctx.Response.AppendBody([]byte(uuid))
}

func (r *Router) deleteTask(ctx *fasthttp.RequestCtx) {
	var (
		uuid string
		ok   bool
	)
	if uuid, ok = ctx.UserValue("uuid").(string); !ok {
		r.logger.Println(ctx.UserValue("uuid"))
		ctx.SetStatusCode(500)
		return
	}
	err := r.service.DeleteTask(ctx, uuid)
	if err != nil {
		r.logger.Println(err)
		ctx.SetStatusCode(500)
		return
	}
}

func (r *Router) startTask(ctx *fasthttp.RequestCtx) {
	var (
		uuid string
		ok   bool
	)
	if uuid, ok = ctx.UserValue("uuid").(string); !ok {
		r.logger.Println(ctx.UserValue("uuid"))
		ctx.SetStatusCode(500)
		return
	}
	err := r.service.StartTask(ctx, uuid)
	if err != nil {
		r.logger.Println(err)
		ctx.SetStatusCode(500)
		return
	}
}
