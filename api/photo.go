package api

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/zi-yang-zhang/audible-silence-backend/service"
)

func getPhotoListHandler(service service.PhotoService, ctx iris.Context) {
	result, err := service.GetPhotoList()
	if err != nil {
		panic(fmt.Errorf("Error finding photo list %s", err))
	}
	formatSuccessResponse(result, ctx)

}

func getPhotoHandler(id uint, service service.PhotoService, ctx iris.Context) {
	result, err := service.GetPhoto(id)
	if err != nil {
		formatSuccessResponse(nil, ctx)
	} else {
		formatSuccessResponse(result, ctx)
	}
}
