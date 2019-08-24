package api

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/zi-yang-zhang/audible-silence-backend/service"
)

type CreateMarkdownRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func initMarkdownAPI() {
	hero.Register(func(ctx iris.Context) (request CreateMarkdownRequest) {
		ctx.ReadJSON(&request)
		return
	})
}
func getMarkdownByIDHandler(id uint, service service.MarkdownService, ctx iris.Context) {
	result, err := service.GetMarkdownByID(id)
	if err != nil {
		formatSuccessResponse(nil, ctx)
	} else {
		formatSuccessResponse(result, ctx)
	}
}

func saveMarkdownHandler(request CreateMarkdownRequest, service service.MarkdownService, ctx iris.Context) {
	err := service.SaveMarkdown(request.Title, request.Content)
	if err != nil {
		formatSuccessResponse(nil, ctx)
	} else {
		formatSuccessResponse(nil, ctx)
	}
}
