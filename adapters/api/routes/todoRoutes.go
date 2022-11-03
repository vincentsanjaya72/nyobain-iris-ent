package routes

import (
	"github.com/kataras/iris/v12"
	todoController "github.com/vcnt72/nyobain-iris-ent/adapters/controller/todo_controller"
)

func todoRoutes(router iris.Party) {
	router.Get("/todos", todoController.Get())
}
