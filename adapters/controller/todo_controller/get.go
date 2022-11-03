package todoController

import (
	"github.com/kataras/iris/v12"
	"github.com/vcnt72/nyobain-iris-ent/shared/response"
)

type GetTodoResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func Get() iris.Handler {
	return func(ctx iris.Context) {
		ctx.JSON(response.Wrap(
			GetTodoResponse{
				Name:        "Ohayou gozaimasu",
				Description: "Mengucapkan selamat pagi kepada nahida~chan",
			},
			nil,
		))
	}
}
