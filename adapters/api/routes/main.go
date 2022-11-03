package routes

import (
	"github.com/kataras/iris/v12"
)

type DummyResponse struct {
	Hello string `json:"hello_world"`
}

func V1(app *iris.Application) iris.Party {
	v1 := app.Party("/v1")
	{
		todoRoutes(v1)
	}

	return v1
}
