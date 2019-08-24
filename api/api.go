package api

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/spf13/viper"
	"github.com/zi-yang-zhang/audible-silence-backend/service"
)

type wrappedContext struct {
	iris.Context
}

//InitAPI initilizes endpoints
func InitAPI(app *iris.Application, db *gorm.DB, serverConfig *viper.Viper) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})
	initMarkdownAPI()
	service.InitPhotoService(db, &service.PhotoLocationConfig{}, app.Logger())
	service.InitMarkdownService(db, app.Logger())
	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions) // <- important for the preflight.
	{
		v1.Get("/photoList", hero.Handler(getPhotoListHandler))
		v1.Get("/photo/{id:uint}", hero.Handler(getPhotoHandler))
		v1.Get("/md/{id:uint}", hero.Handler(getMarkdownByIDHandler))
		v1.Put("/md", hero.Handler(saveMarkdownHandler))
	}
}

func formatResponse(code string, msg string, resp interface{}, ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code": code,
		"msg":  msg,
		"data": resp,
	})
}

func formatSuccessResponse(resp interface{}, ctx iris.Context) {
	formatResponse("0", "success", resp, ctx)
}
