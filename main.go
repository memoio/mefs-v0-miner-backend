package main

import (
	// "encoding/json"
	// "fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"longchain.com/memoriae/orePool/config"
	"longchain.com/memoriae/orePool/web/controllers"
)

func main() {
	addr := config.Addr

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	controllers.New(app)
	app.Run(iris.Addr(addr))
}