package api

import "github.com/valyala/fasthttp"

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
}
