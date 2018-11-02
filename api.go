package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

//InitAPI initilizes endpoints
func initAPI(app *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})
	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions) // <- important for the preflight.
	{
		v1.Get("/ping", createPostHandler())
	}
}

func createPostHandler() context.Handler {
	return func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	}
}
