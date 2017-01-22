package routers

import (
	"ComplimentEngine/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/data", &controllers.DataController{})
	beego.Router("/upload", &controllers.UploadController{})
}
