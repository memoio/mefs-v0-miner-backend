package controllers
import(
	// "fmt"

	"github.com/kataras/iris"

	rm "longchain.com/memoriae/orePool/web/resultModel"
	"longchain.com/memoriae/orePool/web/service"
)

func New(app *iris.Application) {
	app.Get("/list", func(ctx iris.Context){
		ctx.JSON(rm.Ok(service.GetList()))
	})
}