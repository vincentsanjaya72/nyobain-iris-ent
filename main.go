package main

import (
	"github.com/kataras/iris/v12"
	"github.com/vcnt72/nyobain-iris-ent/adapters/api/routes"
)

func main() {

	app := iris.Default()

	routes.V1(app)

	app.Listen(":8080")
}
